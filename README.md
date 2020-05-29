#

/gero --config /usr/local/bin/geroConfig.toml --datadir /root/.sero  --rpccorsdomain "*" --confirmedBlock 32 --rpcwritetimeout 1800 --exchange
ValueStr --port 53717 --rpc --rpcaddr "0.0.0.0" -rpcapi "sero,net,exchange,stake" --rpcport 8545 --exchange --mineMode  $@ 2>> /log/${logName}


#
/gero --config /usr/local/bin/geroConfig.toml --datadir /root/.sero attach



# Go Sero

The SERO system is the world's first blockchain platform to support conﬁdential transactions based on Zero-knowledge proof and supports Turing complete smart contracts.

## What's SERO

https://wiki.sero.cash/en/index.html?file=home-Home


## From source code -- base on CentOS7

https://wiki.sero.cash/en/index.html?file=Start/from-the-sourcecode-base-on-centos7


## How to mine using gero

https://wiki.sero.cash/en/index.html?file=Start/from-the-binary-package

 **Now You don't need license anymore to mine SERO.**


## How to use snapshots

https://wiki.sero.cash/en/index.html?file=Start/sero-chain-snapshot-list


## How to start SERO Staking on gero

https://wiki.sero.cash/en/index.html?file=Tutorial/how-to-staking-using-gero

## Decentralized light wallet

**PC:**
https://github.com/sero-cash/pullup

**Mobile:**
https://sero.cash/app/popup.html


## POW pool list

https://wiki.sero.cash/en/index.html?file=Start/sero-pool-list

## PoS stake node List

https://wiki.sero.cash/en/index.html?file=Start/stake-node-list

## Contribution

Thank you for considering to help out with the source code! We welcome contributions from
anyone on the internet, and are grateful for even the smallest of fixes!

If you'd like to contribute to go-sero, please fork, fix, commit and send a pull request
for the maintainers to review and merge into the main code base. 

Please make sure your contributions adhere to our coding guidelines:

 * Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
 * Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
 * Pull requests need to be based on and opened against the `master` branch.
 * Commit messages should be prefixed with the package(s) they modify.
   * E.g. "sero, rpc: make trace configs optional"

Please see the [Developers' Guide](https://github.com/sero-cash/go-sero/wiki/Developers'-Guide)
for more details on configuring your environment, managing project dependencies and testing procedures.

## Community resources

**Wechat:**  SERO9413

**Discord:**  <https://discord.gg/n5HVxE>

**Twitter:**  <https://twitter.com/SEROdotCASH>

**Telegram:**  <https://t.me/SeroOfficial>

**Gitter:**  <https://gitter.im/sero-cash/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link>


## Other resources

**Official Website:** <https://sero.cash>

**White Paper:** <http://sero-media.s3-website-ap-southeast-1.amazonaws.com/Sero_ENG_V1.06.pdf>

**WIKI:** <https://wiki.sero.cash/zh/index.html?file=home-Home>

**Block Explorer:** <https://explorer.web.sero.cash/blocks.html>

**Introduction Video:** <https://v.qq.com/x/page/s0792e921ok.html>


## License

The go-sero library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The go-sero binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.

*Note: Go Sero inherit with licenses of ethereum.*
