package chaos

import (
	"context"
	tf "github.com/filecoin-project/venus/pkg/testhelpers/testflags"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	mock2 "github.com/filecoin-project/specs-actors/v2/support/mock"
	atesting2 "github.com/filecoin-project/specs-actors/v2/support/testing"
)

func TestSingleton(t *testing.T) {
	tf.UnitTest(t)
	receiver := atesting2.NewIDAddr(t, 100)
	builder := mock2.NewBuilder(context.Background(), receiver)

	rt := builder.Build(t)
	var a Actor

	msg := "constructor should not be called; the Chaos actor is a singleton actor"
	rt.ExpectAssertionFailure(msg, func() {
		rt.Call(a.Constructor, abi.Empty)
	})
	rt.Verify()
}

func TestCallerValidationNone(t *testing.T) {
	tf.UnitTest(t)
	receiver := atesting2.NewIDAddr(t, 100)
	builder := mock2.NewBuilder(context.Background(), receiver)

	rt := builder.Build(t)
	var a Actor

	rt.Call(a.CallerValidation, &CallerValidationArgs{Branch: CallerValidationBranchNone})
	rt.Verify()
}

func TestCallerValidationIs(t *testing.T) {
	tf.UnitTest(t)
	caller := atesting2.NewIDAddr(t, 100)
	receiver := atesting2.NewIDAddr(t, 101)
	builder := mock2.NewBuilder(context.Background(), receiver)

	rt := builder.Build(t)
	rt.SetCaller(caller, builtin2.AccountActorCodeID)
	var a Actor

	caddrs := []address.Address{atesting2.NewIDAddr(t, 101)}

	rt.ExpectValidateCallerAddr(caddrs...)
	// fixed in: https://github.com/filecoin-project/specs-actors/pull/1155
	rt.ExpectAbort(exitcode.SysErrForbidden, func() {
		rt.Call(a.CallerValidation, &CallerValidationArgs{
			Branch: CallerValidationBranchIsAddress,
			Addrs:  caddrs,
		})
	})
	rt.Verify()

	rt.ExpectValidateCallerAddr(caller)
	rt.Call(a.CallerValidation, &CallerValidationArgs{
		Branch: CallerValidationBranchIsAddress,
		Addrs:  []address.Address{caller},
	})
	rt.Verify()
}

func TestCallerValidationType(t *testing.T) {
	tf.UnitTest(t)
	caller := atesting2.NewIDAddr(t, 100)
	receiver := atesting2.NewIDAddr(t, 101)
	builder := mock2.NewBuilder(context.Background(), receiver)

	rt := builder.Build(t)
	rt.SetCaller(caller, builtin2.AccountActorCodeID)
	var a Actor

	rt.ExpectValidateCallerType(builtin2.CronActorCodeID)
	rt.ExpectAbort(exitcode.SysErrForbidden, func() {
		rt.Call(a.CallerValidation, &CallerValidationArgs{
			Branch: CallerValidationBranchIsType,
			Types:  []cid.Cid{builtin2.CronActorCodeID},
		})
	})
	rt.Verify()

	rt.ExpectValidateCallerType(builtin2.AccountActorCodeID)
	rt.Call(a.CallerValidation, &CallerValidationArgs{
		Branch: CallerValidationBranchIsType,
		Types:  []cid.Cid{builtin2.AccountActorCodeID},
	})
	rt.Verify()
}

func TestCallerValidationInvalidBranch(t *testing.T) {
	tf.UnitTest(t)
	receiver := atesting2.NewIDAddr(t, 100)
	builder := mock2.NewBuilder(context.Background(), receiver)

	rt := builder.Build(t)
	var a Actor

	rt.ExpectAssertionFailure("invalid branch passed to CallerValidation", func() {
		rt.Call(a.CallerValidation, &CallerValidationArgs{Branch: -1})
	})
	rt.Verify()
}

func TestDeleteActor(t *testing.T) {
	tf.UnitTest(t)
	receiver := atesting2.NewIDAddr(t, 100)
	beneficiary := atesting2.NewIDAddr(t, 101)
	builder := mock2.NewBuilder(context.Background(), receiver)

	rt := builder.Build(t)
	var a Actor

	rt.ExpectValidateCallerAny()
	rt.ExpectDeleteActor(beneficiary)
	rt.Call(a.DeleteActor, &beneficiary)
	rt.Verify()
}

func TestMutateStateInTransaction(t *testing.T) {
	tf.UnitTest(t)
	receiver := atesting2.NewIDAddr(t, 100)
	builder := mock2.NewBuilder(context.Background(), receiver)

	rt := builder.Build(t)
	var a Actor

	rt.ExpectValidateCallerAny()
	rt.StateCreate(&State{})

	val := "__mutstat test"
	rt.Call(a.MutateState, &MutateStateArgs{
		Value:  val,
		Branch: MutateInTransaction,
	})

	var st State
	rt.GetState(&st)

	if st.Value != val {
		t.Fatal("state was not updated")
	}

	rt.Verify()
}

func TestMutateStateAfterTransaction(t *testing.T) {
	tf.UnitTest(t)
	receiver := atesting2.NewIDAddr(t, 100)
	builder := mock2.NewBuilder(context.Background(), receiver)

	rt := builder.Build(t)
	var a Actor

	rt.ExpectValidateCallerAny()
	rt.StateCreate(&State{})

	val := "__mutstat test"
	rt.Call(a.MutateState, &MutateStateArgs{
		Value:  val,
		Branch: MutateAfterTransaction,
	})

	var st State
	rt.GetState(&st)

	// state should be updated successfully _in_ the transaction but not outside
	if st.Value != val+"-in" {
		t.Fatal("state was not updated")
	}

	rt.Verify()
}

func TestMutateStateReadonly(t *testing.T) {
	tf.UnitTest(t)
	receiver := atesting2.NewIDAddr(t, 100)
	builder := mock2.NewBuilder(context.Background(), receiver)

	rt := builder.Build(t)
	var a Actor

	rt.ExpectValidateCallerAny()
	rt.StateCreate(&State{})

	val := "__mutstat test"
	rt.Call(a.MutateState, &MutateStateArgs{
		Value:  val,
		Branch: MutateReadonly,
	})

	var st State
	rt.GetState(&st)

	if st.Value != "" {
		t.Fatal("state was not expected to be updated")
	}

	rt.Verify()
}

func TestMutateStateInvalidBranch(t *testing.T) {
	tf.UnitTest(t)
	receiver := atesting2.NewIDAddr(t, 100)
	builder := mock2.NewBuilder(context.Background(), receiver)

	rt := builder.Build(t)
	var a Actor

	rt.ExpectValidateCallerAny()
	rt.ExpectAssertionFailure("unknown mutation type", func() {
		rt.Call(a.MutateState, &MutateStateArgs{Branch: -1})
	})
	rt.Verify()
}

func TestAbortWith(t *testing.T) {
	tf.UnitTest(t)
	receiver := atesting2.NewIDAddr(t, 100)
	builder := mock2.NewBuilder(context.Background(), receiver)

	rt := builder.Build(t)
	var a Actor

	msg := "__test forbidden"
	rt.ExpectAbortContainsMessage(exitcode.ErrForbidden, msg, func() {
		rt.Call(a.AbortWith, &AbortWithArgs{
			Code:         exitcode.ErrForbidden,
			Message:      msg,
			Uncontrolled: false,
		})
	})
	rt.Verify()
}

func TestAbortWithUncontrolled(t *testing.T) {
	tf.UnitTest(t)
	receiver := atesting2.NewIDAddr(t, 100)
	builder := mock2.NewBuilder(context.Background(), receiver)

	rt := builder.Build(t)
	var a Actor

	msg := "__test uncontrolled panic"
	rt.ExpectAssertionFailure(msg, func() {
		rt.Call(a.AbortWith, &AbortWithArgs{
			Message:      msg,
			Uncontrolled: true,
		})
	})
	rt.Verify()
}

func TestInspectRuntime(t *testing.T) {
	tf.UnitTest(t)
	caller := atesting2.NewIDAddr(t, 100)
	receiver := atesting2.NewIDAddr(t, 101)
	builder := mock2.NewBuilder(context.Background(), receiver)

	rt := builder.Build(t)
	rt.SetCaller(caller, builtin2.AccountActorCodeID)
	rt.StateCreate(&State{})
	var a Actor

	rt.ExpectValidateCallerAny()
	ret := rt.Call(a.InspectRuntime, abi.Empty)
	rtr, ok := ret.(*InspectRuntimeReturn)
	if !ok {
		t.Fatal("invalid return value")
	}
	if rtr.Caller != caller {
		t.Fatal("unexpected runtime caller")
	}
	if rtr.Receiver != receiver {
		t.Fatal("unexpected runtime receiver")
	}
	rt.Verify()
}
