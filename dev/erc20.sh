res=$(fbchaincli tx wasm store ./wasm/erc20/artifacts/cw_erc20-aarch64.wasm --fees 0.01fibo --from captain --gas=2000000 -b block -y)
code_id=$(echo "$res" | jq '.logs[0].events[1].attributes[0].value' | sed 's/\"//g')
res=$(fbchaincli tx wasm instantiate "$code_id" '{"decimals":10,"initial_balances":[{"address":"0xbbE4733d85bc2b90682147779DA49caB38C0aA1F","amount":"100000000"}],"name":"my test token", "symbol":"MTT"}' --label test1 --admin fb1h0j8x0v9hs4eq6ppgamemfyu4vuvp2sl0q9p3v --fees 0.001fbio --from captain -b block -y)
contractAddr=$(echo "$res" | jq '.logs[0].events[0].attributes[0].value' | sed 's/\"//g')
fbchaincli tx wasm execute "$contractAddr" '{"transfer":{"amount":"100","recipient":"0xCf164e001d86639231d92Ab1D71DB8353E43C295"}}' --fees 0.001fibo --from captain -b block -y

echo " ========================================================== "
echo "## show all codes uploaded ##"
fbchaincli query wasm list-code

echo " ========================================================== "
echo "## show contract info by contract addr ##"
fbchaincli query wasm contract "$contractAddr"

echo " ========================================================== "
echo "## show contract update history by contract addr ##"
fbchaincli query wasm contract-history "$contractAddr"

echo " ========================================================== "
echo "## query contract state by contract addr ##"
echo "#### all state"
fbchaincli query wasm contract-state all "$contractAddr"
echo "#### raw state"
fbchaincli query wasm contract-state raw "$contractAddr" 0006636F6E666967636F6E7374616E7473
echo "#### smart state"
echo "$contractAddr"
fbchaincli query wasm contract-state smart "$contractAddr" '{"balance":{"address":"0xbbE4733d85bc2b90682147779DA49caB38C0aA1F"}}'
fbchaincli query wasm contract-state smart "$contractAddr" '{"balance":{"address":"0xCf164e001d86639231d92Ab1D71DB8353E43C295"}}'



