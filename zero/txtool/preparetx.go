package txtool

/*
import (
	"fmt"
	"github.com/sero-cash/go-czero-import/keys"
	"github.com/sero-cash/go-sero/zero/txs/assets"
	"github.com/sero-cash/go-sero/zero/utils"
	"math/big"
	"strings"
)

type Reception struct {
	Addr     keys.PKr
	Currency string
	Value    *big.Int
}

type PreTxParam struct {
	From       keys.Uint512
	RefundTo   *keys.PKr
	Receptions []Reception
	Gas        uint64
	GasPrice   *big.Int
	Roots      []keys.Uint256
}

func GenTxParam(param PreTxParam) (txParam *txtool.GTxParam, e error) {
	utxos, err := preGenTx(param)
	if err != nil {
		return nil, err
	}

	if value, ok := self.accounts.Load(param.From); ok {
		var refundTo keys.PKr
		if param.RefundTo == nil {
			account := value.(*Account)
			refundTo = account.mainPkr
		} else {
			refundTo = *param.RefundTo
		}
		txParam, e = buildTxParam(utxos, &refundTo, param.Receptions, param.Gas, param.GasPrice)
	} else {
		return nil, errors.New("not found Pk")
	}

	return
}

func preGenTx(param PreTxParam) (utxos []keys.Uint256, err error) {
	if len(param.Roots) > 0 {
		utxos = param.Roots
	} else {
		amounts := map[string]*big.Int{}
		for _, each := range param.Receptions {
			if amount, ok := amounts[each.Currency]; ok {
				amount.Add(amount, each.Value)
			} else {
				amounts[each.Currency] = new(big.Int).Set(each.Value)
			}
		}
		if amount, ok := amounts["SERO"]; ok {
			amount.Add(amount, new(big.Int).Mul(new(big.Int).SetUint64(param.Gas), param.GasPrice))
		} else {
			amounts["SERO"] = new(big.Int).Mul(new(big.Int).SetUint64(param.Gas), param.GasPrice)
		}
		for currency, amount := range amounts {
			list, remain := self.findUtxos(&param.From, currency, amount)
			if remain.Sign() > 0 {
				return utxos, errors.New(fmt.Sprintf("not enough token, maximum available token is %s", new(big.Int).Sub(amount, remain).String()))
			} else {
				utxos = append(utxos, list...)
			}
		}
	}
	count := 0
	for _, each := range utxos {
		if !each.IsZ {
			count++
		}
	}
	if count > 2500 {
		err = errors.New("ins.len > 2500")
	}
	return
}

func buildTxParam(
	utxos []keys.Uint256,
	refundTo *keys.PKr,
	receptions []Reception,
	gas uint64,
	gasPrice *big.Int) (txParam *txtool.GTxParam, e error) {

	txParam = new(txtool.GTxParam)
	txParam.Gas = gas
	txParam.GasPrice = *gasPrice

	txParam.From = txtool.Kr{PKr: *refundTo}

	Ins := []txtool.GIn{}
	wits, err := txtool.SRI_Inst.GetAnchor(utxos)
	if err != nil {
		e = err
		return
	}

	amounts := make(map[string]*big.Int)
	ticekts := make(map[keys.Uint256]keys.Uint256)
	for index, utxo := range utxos {
		if out := txtool.GetOut(&utxo, 0); out != nil {
			Ins = append(Ins, txtool.GIn{Out: txtool.Out{Root: utxo, State: *out}, Witness: wits[index]})

			if out.OS.Out_O.Asset.Tkn != nil {
				currency := strings.Trim(string(out.OS.Out_O.Asset.Tkn.Currency[:]), string([]byte{0}))
				if amount, ok := amounts[currency]; ok {
					amount.Add(amount, out.OS.Out_O.Asset.Tkn.Value.ToIntRef())
				} else {
					amounts[currency] = new(big.Int).Set(out.OS.Out_O.Asset.Tkn.Value.ToIntRef())
				}

			}
			if out.OS.Out_O.Asset.Tkt != nil {
				ticekts[out.OS.Out_O.Asset.Tkt.Value] = out.OS.Out_O.Asset.Tkt.Category
			}
		}
	}

	Outs := []txtool.GOut{}
	for _, reception := range receptions {
		currency := strings.ToUpper(reception.Currency)
		if amount, ok := amounts[currency]; ok && amount.Cmp(reception.Value) >= 0 {

			if self.isPk(reception.Addr) {
				pk := reception.Addr.ToUint512()
				pkr := self.createPkr(&pk, 1)
				Outs = append(Outs, txtool.GOut{PKr: pkr, Asset: assets.Asset{Tkn: &assets.Token{
					Currency: *common.BytesToHash(common.LeftPadBytes([]byte(currency), 32)).HashToUint256(),
					Value:    utils.U256(*reception.Value),
				}}})
			} else {
				Outs = append(Outs, txtool.GOut{PKr: reception.Addr, Asset: assets.Asset{Tkn: &assets.Token{
					Currency: *common.BytesToHash(common.LeftPadBytes([]byte(currency), 32)).HashToUint256(),
					Value:    utils.U256(*reception.Value),
				}}})
			}

			amount.Sub(amount, reception.Value)
			if amount.Sign() == 0 {
				delete(amounts, currency)
			}
		}

	}

	fee := new(big.Int).Mul(new(big.Int).SetUint64(gas), gasPrice)
	if amount, ok := amounts["SERO"]; !ok || amount.Cmp(fee) < 0 {
		e = fmt.Errorf("Exchange Error: not enough")
		return
	} else {
		amount.Sub(amount, fee)
		if amount.Sign() == 0 {
			delete(amounts, "SERO")
		}
	}

	if len(amounts) > 0 {
		for currency, value := range amounts {
			Outs = append(Outs, txtool.GOut{PKr: txParam.From.PKr, Asset: assets.Asset{Tkn: &assets.Token{
				Currency: *common.BytesToHash(common.LeftPadBytes([]byte(currency), 32)).HashToUint256(),
				Value:    utils.U256(*value),
			}}})
		}
	}
	if len(ticekts) > 0 {
		for value, category := range ticekts {
			Outs = append(Outs, txtool.GOut{PKr: txParam.From.PKr, Asset: assets.Asset{Tkt: &assets.Ticket{
				Category: category,
				Value:    value,
			}}})
		}
	}

	txParam.Ins = Ins
	txParam.Outs = Outs

	for _, utxo := range utxos {
		self.usedFlag.Store(utxo.Root, 1)
	}

	return
}
*/
