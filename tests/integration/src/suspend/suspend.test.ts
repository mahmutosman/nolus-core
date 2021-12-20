import {MsgChangeSuspended, protobufPackage} from "../util/codec/nolus/suspend/v1beta1/tx";
import {SigningCosmWasmClient} from "@cosmjs/cosmwasm-stargate";
import {
    DEFAULT_FEE,
    getGenesisUser1Client,
    getGenesisUser1Wallet,
    getUser1Wallet,
    getValidatorClient,
    getValidatorWallet
} from "../util/clients";
import Long from "long";
import {EncodeObject} from "@cosmjs/proto-signing/build/registry";
import {getSuspendQueryClient} from "./suspend-client";
import {QueryClientImpl} from "../util/codec/nolus/suspend/v1beta1/query";
import {AccountData} from "@cosmjs/proto-signing";
import {isBroadcastTxFailure} from "@cosmjs/stargate";


describe("suspend module", () => {
    const DUMMY_TRANFER_MSG = [{denom: "unolus", "amount": "1"}];
    let validatorClient: SigningCosmWasmClient;
    let validatorAccount: AccountData
    let genUserClient: SigningCosmWasmClient;
    let genUserAccount: AccountData
    let suspendQueryClient: QueryClientImpl

    beforeEach(async () => {
        validatorClient = await getValidatorClient();
        [validatorAccount] = await (await getValidatorWallet()).getAccounts();
        genUserClient = await getGenesisUser1Client();
        [genUserAccount] = await (await getGenesisUser1Wallet()).getAccounts()
        suspendQueryClient = await getSuspendQueryClient(process.env.NODE_URL as string);
    })

    afterAll(async () => {
        const suspendedMsg = asMsgChangeSuspended(validatorAccount, false, 0);
        await validatorClient.signAndBroadcast(validatorAccount.address, [suspendedMsg], DEFAULT_FEE);
    })

    describe("given admin account signer", () => {
        test("suspended state can be enabled and then disabled", async () => {
            const suspendedMsg = asMsgChangeSuspended(validatorAccount, true, 1);
            await validatorClient.signAndBroadcast(validatorAccount.address, [suspendedMsg], DEFAULT_FEE);

            let suspendResponse = await suspendQueryClient.SuspendedState({});
            expect(suspendResponse.state).toBeDefined();
            expect(suspendResponse.state?.suspended).toBeTruthy();
            expect(suspendResponse.state?.blockHeight).toEqual(Long.fromNumber(1));

            const unsuspendedMsg = asMsgChangeSuspended(validatorAccount, false, 2);
            await validatorClient.signAndBroadcast(validatorAccount.address, [unsuspendedMsg], DEFAULT_FEE);

            suspendResponse = await suspendQueryClient.SuspendedState({});
            expect(suspendResponse.state).toBeDefined();
            expect(suspendResponse.state?.suspended).toBeFalsy();
            expect(suspendResponse.state?.blockHeight).toEqual(Long.fromNumber(2));
        })

        test("when suspended and height reached then other transactions are unauthorized", async () => {
            let currentHeight = await validatorClient.getHeight();
            const suspendedMsg = asMsgChangeSuspended(validatorAccount, true, currentHeight + 1);
            await validatorClient.signAndBroadcast(validatorAccount.address, [suspendedMsg], DEFAULT_FEE);

            const broadcast = () => validatorClient.sendTokens(validatorAccount.address, genUserAccount.address, DUMMY_TRANFER_MSG, DEFAULT_FEE)
            await expect(broadcast).rejects.toThrow(/^.*node is suspended: unauthorized$/)
        })

        test("when suspended but height is not reached then other transactions are authorized", async () => {
            let currentHeight = await validatorClient.getHeight();
            const suspendedMsg = asMsgChangeSuspended(validatorAccount, true, currentHeight + 1000);
            await validatorClient.signAndBroadcast(validatorAccount.address, [suspendedMsg], DEFAULT_FEE);

            const result = await validatorClient.sendTokens(validatorAccount.address, genUserAccount.address, DUMMY_TRANFER_MSG, DEFAULT_FEE)
            expect(isBroadcastTxFailure(result)).toBeFalsy()
        })

        test("when suspended can send multiple messages when one of them is MsgChangeSuspend", async () => {
            let currentHeight = await validatorClient.getHeight();
            const suspendedMsg = asMsgChangeSuspended(validatorAccount, true, currentHeight);
            await validatorClient.signAndBroadcast(validatorAccount.address, [suspendedMsg], DEFAULT_FEE);

            const unsuspendMsg = asMsgChangeSuspended(validatorAccount, false, 0);
            const [user1Account] = await (await getUser1Wallet()).getAccounts()

            const fullTransferMsg1 = dummyMsgTransfer(validatorAccount, genUserAccount);
            const fullTransferMsg2 = dummyMsgTransfer(validatorAccount, user1Account);

            const result = await validatorClient.signAndBroadcast(validatorAccount.address, [fullTransferMsg1, unsuspendMsg, fullTransferMsg2], DEFAULT_FEE);
            expect(isBroadcastTxFailure(result)).toBeFalsy();

            const suspendResponse = await suspendQueryClient.SuspendedState({});
            expect(suspendResponse.state?.suspended).toBeFalsy();
        })
    });

    describe("given non admin account signer", () => {
        test("state suspended cannot be changed", async () => {
            // ensure state cannot be modified while being in state suspended
            const suspendedMsg = asMsgChangeSuspended(validatorAccount, true, 0);
            await validatorClient.signAndBroadcast(validatorAccount.address, [suspendedMsg], DEFAULT_FEE);

            let invalidMsg = asMsgChangeSuspended(genUserAccount, false, 1);
            let result = await genUserClient.signAndBroadcast(genUserAccount.address, [invalidMsg], DEFAULT_FEE)
            expect(isBroadcastTxFailure(result)).toBeTruthy();

            // ensure state cannot be modified while also being in state unsuspended
            const unSuspendedMsg = asMsgChangeSuspended(validatorAccount, false, 0);
            await validatorClient.signAndBroadcast(validatorAccount.address, [unSuspendedMsg], DEFAULT_FEE);


            invalidMsg = asMsgChangeSuspended(genUserAccount, false, 1);
            result = await genUserClient.signAndBroadcast(genUserAccount.address, [invalidMsg], DEFAULT_FEE)
            expect(isBroadcastTxFailure(result)).toBeTruthy();
        })

        test("state height cannot be changed", async () => {
            let invalidMsg = asMsgChangeSuspended(genUserAccount, false, 100);
            let result = await genUserClient.signAndBroadcast(genUserAccount.address, [invalidMsg], DEFAULT_FEE)
            expect(isBroadcastTxFailure(result)).toBeTruthy();
        })

        test("message cannot be send with forged from address", async () => {
            let invalidMsg = asMsgChangeSuspended(validatorAccount, false, 100);
            let broadcast = () => genUserClient.signAndBroadcast(genUserAccount.address, [invalidMsg], DEFAULT_FEE);
            await expect(broadcast).rejects.toThrow(/^Broadcasting transaction failed with code 8*/)
        })

        test("when suspended and height reached then other transactions are unauthorized", async () => {
            const suspendedMsg = asMsgChangeSuspended(validatorAccount, true, 0);
            await validatorClient.signAndBroadcast(validatorAccount.address, [suspendedMsg], DEFAULT_FEE);

            const broadcast = () => genUserClient.sendTokens(genUserAccount.address, validatorAccount.address, DUMMY_TRANFER_MSG, DEFAULT_FEE)
            await expect(broadcast).rejects.toThrow(/^.*node is suspended: unauthorized$/)
        })

        test("when suspended cannot bypass it by sending multiple messages including MsgChangeSuspend", async () => {
            const suspendedMsg = asMsgChangeSuspended(validatorAccount, true, 0);
            await validatorClient.signAndBroadcast(validatorAccount.address, [suspendedMsg], DEFAULT_FEE);

            const unsuspendMsg = asMsgChangeSuspended(genUserAccount, false, 0);
            const [user1Account] = await (await getUser1Wallet()).getAccounts()

            const fullTransferMsg1 = dummyMsgTransfer(genUserAccount, validatorAccount)
            const fullTransferMsg2 = dummyMsgTransfer(genUserAccount, user1Account)

            const result = await genUserClient.signAndBroadcast(genUserAccount.address, [fullTransferMsg1, unsuspendMsg, fullTransferMsg2], DEFAULT_FEE);
            console.log(result)
            expect(isBroadcastTxFailure(result)).toBeTruthy();
        })
    })

    function asMsgChangeSuspended(fromAddress: AccountData, suspended: boolean, blockHeight: number): EncodeObject {
        const suspendMsg: MsgChangeSuspended = {
            fromAddress: fromAddress.address,
            suspended,
            blockHeight: Long.fromNumber(blockHeight)
        }
        return {
            typeUrl: `/${protobufPackage}.MsgChangeSuspended`,
            value: suspendMsg,
        };
    }

    function dummyMsgTransfer(fromAddress: AccountData, toAddress: AccountData): EncodeObject {
        return {
            typeUrl: "/cosmos.bank.v1beta1.MsgSend",
            value: {
                fromAddress: fromAddress.address,
                toAddress: toAddress.address,
                amount: DUMMY_TRANFER_MSG
            }
        };
    }
})
