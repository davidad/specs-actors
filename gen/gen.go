package main

import (
	gen "github.com/whyrusleeping/cbor-gen"

	abi "github.com/filecoin-project/specs-actors/actors/abi"
	builtin "github.com/filecoin-project/specs-actors/actors/builtin"
	account "github.com/filecoin-project/specs-actors/actors/builtin/account"
	cron "github.com/filecoin-project/specs-actors/actors/builtin/cron"
	init_ "github.com/filecoin-project/specs-actors/actors/builtin/init"
	market "github.com/filecoin-project/specs-actors/actors/builtin/market"
	miner "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	multisig "github.com/filecoin-project/specs-actors/actors/builtin/multisig"
	paych "github.com/filecoin-project/specs-actors/actors/builtin/paych"
	power "github.com/filecoin-project/specs-actors/actors/builtin/power"
	reward "github.com/filecoin-project/specs-actors/actors/builtin/reward"
)

func main() {
	// Common types
	if err := gen.WriteTupleEncodersToFile("./actors/abi/cbor_gen.go", "abi",
		abi.PieceInfo{},
		abi.SectorID{},
		abi.SealVerifyInfo{},
		abi.OnChainSealVerifyInfo{},
		abi.PoStCandidate{},
		abi.PoStProof{},
		abi.PrivatePoStCandidateProof{},
		abi.OnChainPoStVerifyInfo{},
		abi.OnChainElectionPoStVerifyInfo{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./actors/builtin/cbor_gen.go", "builtin",
		builtin.MinerAddrs{},
	); err != nil {
		panic(err)
	}

	// Actors
	if err := gen.WriteTupleEncodersToFile("./actors/builtin/account/cbor_gen.go", "account",
		// actor state
		account.State{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./actors/builtin/init/cbor_gen.go", "init",
		// actor state
		init_.State{},
		// method params
		init_.ConstructorParams{},
		init_.ExecParams{},
		init_.ExecReturn{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./actors/builtin/cron/cbor_gen.go", "cron",
		// actor state
		cron.State{},
		cron.Entry{},
		// method params
		cron.ConstructorParams{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./actors/builtin/reward/cbor_gen.go", "reward",
		// actor state
		reward.State{},
		// method params
		reward.AwardBlockRewardParams{},
		// other types
		reward.Reward{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./actors/builtin/multisig/cbor_gen.go", "multisig",
		// actor state
		multisig.State{},
		multisig.Transaction{},
		// method params
		multisig.ConstructorParams{},
		multisig.ProposeParams{},
		multisig.AddSignerParams{},
		multisig.RemoveSignerParams{},
		multisig.TxnIDParams{},
		multisig.ChangeNumApprovalsThresholdParams{},
		multisig.SwapSignerParams{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./actors/builtin/paych/cbor_gen.go", "paych",
		// actor state
		paych.State{},
		paych.LaneState{},
		paych.Merge{},
		// method params
		paych.ConstructorParams{},
		paych.UpdateChannelStateParams{},
		paych.SignedVoucher{},
		paych.ModVerifyParams{},
		paych.PaymentVerifyParams{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./actors/builtin/power/cbor_gen.go", "power",
		// actors state
		power.State{},
		power.Claim{},
		power.CronEvent{},
		// method params
		power.AddBalanceParams{},
		power.CreateMinerParams{},
		power.DeleteMinerParams{},
		power.WithdrawBalanceParams{},
		power.EnrollCronEventParams{},
		power.OnSectorTerminateParams{},
		power.OnSectorModifyWeightDescParams{},
		power.OnSectorProveCommitParams{},
		power.ReportConsensusFaultParams{},
		power.OnMinerWindowedPoStFailureParams{},
		power.OnSectorTemporaryFaultEffectiveEndParams{},
		power.OnSectorTemporaryFaultEffectiveBeginParams{},
		// method returns
		power.CreateMinerReturn{},
		// other types
		power.MinerConstructorParams{},
		power.SectorStorageWeightDesc{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./actors/builtin/market/cbor_gen.go", "market",
		// actor state
		market.State{},

		// method params
		market.WithdrawBalanceParams{},
		market.PublishStorageDealsParams{},
		market.VerifyDealsOnSectorProveCommitParams{},
		market.ComputeDataCommitmentParams{},
		market.OnMinerSectorsTerminateParams{},
		market.HandleExpiredDealsParams{},
		// method returns
		market.PublishStorageDealsReturn{},
		// other types
		market.DealProposal{},
		market.ClientDealProposal{},
		market.DealState{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./actors/builtin/miner/cbor_gen.go", "miner",
		// actor state
		miner.State{},
		miner.MinerInfo{},
		miner.PoStState{},
		miner.SectorPreCommitOnChainInfo{},
		miner.SectorPreCommitInfo{},
		miner.SectorOnChainInfo{},
		miner.WorkerKeyChange{},
		// method params
		// miner.ConstructorParams{},
		miner.TerminateSectorsParams{},
		miner.ChangePeerIDParams{},
		miner.ProveCommitSectorParams{},
		miner.ChangeWorkerAddressParams{},
		miner.ExtendSectorExpirationParams{},
		miner.DeclareTemporaryFaultsParams{},
		miner.GetControlAddressesReturn{},
		miner.CheckSectorProvenParams{},
		// other types
		miner.CronEventPayload{},
	); err != nil {
		panic(err)
	}

}
