package oracle_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/ethertest"
)

var _ = Describe("updateRateManual", func() {

	Context("When the token is already supported", func() {
		BeforeEach(func() {
			tx, err := oracle.AddTokens(controller.TransactOpts(), []common.Address{common.HexToAddress("0x1")}, stringsToByte32("ETH"), []byte{8})
			Expect(err).ToNot(HaveOccurred())
			be.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())
		})
		Context("When called by the controller", func() {
			It("Should not fail", func() {
				tx, err := oracle.UpdateTokenRate(controller.TransactOpts(), common.HexToAddress("0x1"), big.NewInt(666))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeTrue())
			})
		})
		Context("When not called by the controller", func() {
			It("Should fail", func() {
				tx, err := oracle.UpdateTokenRate(randomAccount.TransactOpts(WithGasLimit(100000)), common.HexToAddress("0x1"), big.NewInt(666))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})

	Context("When the token is not supported", func() {
		Context("When called by the controller", func() {
			It("Should fail", func() {
				tx, err := oracle.UpdateTokenRate(controller.TransactOpts(WithGasLimit(100000)), common.HexToAddress("0x1"), big.NewInt(666))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
		Context("When not called by the controller", func() {
			It("Should fail", func() {
				tx, err := oracle.UpdateTokenRate(randomAccount.TransactOpts(WithGasLimit(100000)), common.HexToAddress("0x1"), big.NewInt(666))
				Expect(err).ToNot(HaveOccurred())
				be.Commit()
				Expect(isSuccessful(tx)).To(BeFalse())
			})
		})
	})

})
