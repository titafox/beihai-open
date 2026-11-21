// import { ethers } from "ethers";

// import ethersJson from "@/utils/purse";
// import { get } from '@/http';


/**
 * 接口获取403 调用登录接口重新获取token 调用第一次授权的签名
 */
let metamask = {}

if (window.ethereum) {
	const isInstall = () => {
		return typeof window.ethereum !== 'undefined' ? true : false;
	}
	// let WALLETSTR = location.href.split('://')[0] === 'http' ? '0xf78431c27ebc2a00f9db9697fe02d89c295bc85f' :
	// 	'0x1999067838b109268252382c0a7d6612b6648eb4';

	const provider = new ethers.providers.Web3Provider(window.ethereum);

	// let contract = new ethers.Contract(
	// 	WALLETSTR,
	// 	ethersJson,
	// 	provider
	// );

	// 获取用户nonce
	const getAuthorNonce = async (data) => await userProfile(data);
	// const getAuthorNonce = async (data: Wallet) => await get('user/profile', data);

	// isApprovedForAll 判断是否授权
	const hasChained = async () => {
		// console.log(window.ethereum.selectedAddress, WALLETSTR)
		const chainId = getCookie('chainId');
		// let type = 'dog'
		// if(chainId == '0x1'){
		// 	type='nft'
		// }
		let res = await getIsApprovalForAllInfo()

		let WALLETSTR = res.contract_address

		let ethersJson = res.abi;


		contract = new ethers.Contract(
			WALLETSTR,
			ethersJson,
			provider
		);
		return await contract.isApprovedForAll(window.ethereum.selectedAddress, WALLETSTR)
	}



	//铸件，导入，创建合集的铸件的时候获取签名
	const createSignMessage = async () => {
		const signer = provider.getSigner();
		const wallet = window.ethereum.selectedAddress; //当前钱包地址
		const userData = await getAuthorNonce({
			wallet
		});
		const msg = "Nonce:" + userData.nonce;
		var data = ''
		try{
			data = await signer.signMessage(msg)
		}catch(e){
			console.log(e)
			showModal({
				message: e.message,
				type: 'error'
			})
		}


		// const data = await signer.signMessage(msg);

		return data;
	}



	//获取签名后如果 上链false setApprovalForAll获取hash 然后getTransactionReceipt监听
	const getStatusWithHash = async (hash, callback) => {

		let info = await provider.getTransactionReceipt(hash);
		if (info && info.status === 1) {
			callback()
		} else {
			// window.queryWorkStatus = setInterval(async () => {
			// 	info = await provider.getTransactionReceipt(hash);
			// 	if (info && info.status === 1) {
			// 		callback()
			// 		clearInterval(window.queryWorkStatus)
			// 	}
			// }, 1000);
			setTimeout( async ()=>{
				await getStatusWithHash(hash,callback)
			},1000)
		}
	}

	//铸件，导入，创建合集的铸件的时候获取钱包授权
	const getHashData = async () => {


		const chainId = getCookie('chainId');
		// let type = 'dog'
		// if(chainId == '0x1'){
		// 	type='nft'
		// }
		let res = await getsetApprovalForAllInfo()

		let WALLETSTR = res.contract_address

		let ethersJson = res.abi;


		let contract = new ethers.Contract(
			WALLETSTR,
			ethersJson,
			provider
		);


		const signer = provider.getSigner();
		const daiWithSigner = contract.connect(signer);
		console.log(daiWithSigner)
		console.log(WALLETSTR)
		const data = await daiWithSigner.setApprovalForAll(WALLETSTR, true);
		console.log(data)
		return data;
	}

	// 接口获取403 调用登录接口重新获取token
	// 钱包鉴权
	const getTokenInfo = async () => {
		const wallet = window.ethereum.selectedAddress || getCookie('accountsKey');
		console.log(wallet)
		if (!wallet) {
			onLinkPurse()
		}
		// 签名获取sign

		console.log('获取token')
		const signData = await loginSignMessage();
		// 获取token
		const data = await userSign({
			sign: signData,
			wallet: wallet, //window.ethereum.selectedAddress
		}).then(res => {
			console.log(res)
			getCookie('token', res.token)
			getCookie('accountsKey',wallet)
			location.reload()
		})

		return data;

	}


	//第一次授权签名
	// 调用第一次授权的签名
	// createUser
	const loginSignMessage = async () => {
		const signer = provider.getSigner();
		const wallet = window.ethereum.selectedAddress || getCookie('accountsKey');

		const userData = await getAuthorNonce({
			wallet
		});
		const msg = "Welcome to Coltstail!\n" +
			"\n" +
			"Click \"Sign\" to sign in. No password needed!\n" +
			"This request will not trigger a blockchain transaction or cost any gas fees.\n" +
			"\n" +
			"Your authentication status will be reset after 24 hours.\n" +
			"\n" +
			"Wallet address:\n" + wallet +
			"\n" +
			"\n" +
			"Nonce:\n" +
			userData.nonce +
			"\n";

		const data = await signer.signMessage(msg);
		return data;
	}

	//购买的时候获取hash
	const sellOrderNft = async ({
									value,
									from,
									to,
									nft_token_id: token_id,
									percentage_basis_points: _percentageBasisPoints,
									token_uri,
								}) => {



		const chainId = getCookie('chainId');
		// let type = 'dog'
		// if(chainId == '0x1'){
		// 	type='nft'
		// }
		let res = await getBuyTokenInfo()

		let WALLETSTR = res.contract_address

		let ethersJson = res.abi;


		let contract = new ethers.Contract(
			WALLETSTR,
			ethersJson,
			provider
		);



		const signer = provider.getSigner()
		const daiWithSigner = contract.connect(signer);



		return await daiWithSigner.buyToken(
			from, to, token_id, token_uri, _percentageBasisPoints, {
				value: value
			}
		);
	}

	//获取账户余额
	const getBalance = async (address) => {
		const res = await provider.getBalance(address);


		let etherString = ethers.utils.formatEther(res);
		return etherString
	}

	// 增加 马蹄链测试网络
	const addPolygonTestnet = async () => {
		try {
			await window.ethereum.request({
				method: 'wallet_switchEthereumChain',
				params: [{
					chainId: '0x13881'
				}],
			});
		} catch (switchError) {
			// This error code indicates that the chain has not been added to MetaMask.
			if (switchError.code === 4902) {
				try {
					await window.ethereum.request({
						method: 'wallet_addEthereumChain',
						params: [{
							chainId: '0x13881',
							chainName: 'Polygon Testnet',
							nativeCurrency: {
								name: 'MATIC',
								symbol: 'MATIC',
								decimals: 18
							},
							rpcUrls: ['https://rpc-mumbai.maticvigil.com/'],
							blockExplorerUrls: ['https://mumbai.polygonscan.com/']
						}]
					});
				} catch (addError) {
					// handle "add" error
				}
			}
		}
	}

	// 增加马蹄链主网络
	const addMaticMainnet = async () => {
		try {
			await window.ethereum.request({
				method: 'wallet_switchEthereumChain',
				params: [{
					chainId: '0x5532c'
				}],
			});
		} catch (switchError) {
			console.log(switchError)
			// This error code indicates that the chain has not been added to MetaMask.
			if (switchError.code === 4902) {
				try {
					await window.ethereum.request({
						method: 'wallet_addEthereumChain',
						params: [{
							chainId: '0x5532c',
							chainName: 'Meta Mainnet',
							nativeCurrency: {
								name: 'Meta',
								symbol: 'Meta',
								decimals: 18
							},
							rpcUrls: ['https://rpc.metaideachain.org/'],
							blockExplorerUrls: ['https://www.metaideachain.org/'],
							iconUrls: ['https://www.metaideachain.org/met_logo2.png'],

						}],
					});
				} catch (addError) {
					// handle "add" error
				}
			}
		}
	}

	//获取账户 CTS 余额
	const getBalanceOfCTS = async (address) => {

		let res = await balanceOfCTS({
			type:'tail'
		})

		let WALLETSTR = res.contract_address;

		// provider = new ethers.providers.Web3Provider(window.ethereum);

		let ethersJson = res.abi

		let contract = new ethers.Contract(
			WALLETSTR,
			ethersJson,
			provider
		);

		const signer = provider.getSigner()
		const daiWithSigner = contract.connect(signer);
		let amount =  await daiWithSigner.balanceOf(
			address
		);

		return ethers.utils.formatEther(amount);
	}


	const ticketApprove = async (data) => {

		let res = await getApprovalInfo({
			type:'tail'
		})

		let WALLETSTR = res.contract_address;
		let spender_address = res.spender_address

		// provider = new ethers.providers.Web3Provider(window.ethereum);

		let ethersJson = res.abi
		console.log(ethersJson)

		let contract = new ethers.Contract(
			WALLETSTR,
			ethersJson,
			provider
		);


		const signer = provider.getSigner()
		const daiWithSigner = contract.connect(signer);
		return await daiWithSigner.approve(
			spender_address,ethers.utils.parseEther(data.price)
		);
	}


	// const ticketisApproved = async (data) => {

	// 	let res = getIsApprovalForAllInfo({
	// 		type:'tail'
	// 	})

	// 	WALLETSTR = res.contract_address;

	// 	// provider = new ethers.providers.Web3Provider(window.ethereum);

	// 	ethersJson = res.abi
	// 	console.log(ethersJson)

	// 	contract = new ethers.Contract(
	// 		WALLETSTR,
	// 		ethersJson,
	// 		provider
	// 	);


	// 	return await contract.isApprovedForAll(window.ethereum.selectedAddress, WALLETSTR)
	// }

	const ticketBuy = async (data) => {
		let res = await getBuyContract({
			type:'tail'
		})
		let WALLETSTR = res.contract_address;

		// provider = new ethers.providers.Web3Provider(window.ethereum);

		let ethersJson = res.abi
		console.log(ethersJson)

		let contract = new ethers.Contract(
			WALLETSTR,
			ethersJson,
			provider
		);

		const signer = provider.getSigner()
		const daiWithSigner = contract.connect(signer);
		//console.log(ethers.BigNumber.from(data.tokenId));
		//let tokenId = "'"+parseInt(data.tokenId)+"'";
		//console.log(tokenId)

		return await daiWithSigner.buy(
			// from, to, token_id, token_uri, _percentageBasisPoints, {
			// 	value: value
			// }

			//ethers.BigNumber.from(1483423941295874048)
			data.tokenId

		);
	}


	const ticketSell = async (data)=> {
		let res = await getTicketSellInfo({
			type: 'tail'
		})

		let WALLETSTR = res.contract_address;

		// provider = new ethers.providers.Web3Provider(window.ethereum);

		let ethersJson = res.abi
		console.log(ethersJson)

		let contract = new ethers.Contract(
			WALLETSTR,
			ethersJson,
			provider
		);

		const signer = provider.getSigner()
		const daiWithSigner = contract.connect(signer);

		return await daiWithSigner.sell(
			data.tokenId,
			data.amount
		);
	}





	metamask = {
		isInstall: isInstall,
		hasChained: hasChained,
		getHashData: getHashData,
		createSignMessage: createSignMessage,
		getStatusWithHash: getStatusWithHash,
		getTokenInfo: getTokenInfo,
		sellOrderNft: sellOrderNft,
		getBalance: getBalance,
		getBalanceOfCTS: getBalanceOfCTS,
		ticketApprove:ticketApprove,
		ticketBuy:ticketBuy,
		ticketSell:ticketSell
		// addPolygonTestnet: addPolygonTestnet,
		// addMaticMainnet: addMaticMainnet

	}
}
