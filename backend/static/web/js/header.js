$(function() {
	//navbar 下划线的判断
	let navIndex = -1;
	if(window.location.pathname.indexOf('index') > -1 || window.location.pathname == '/'){
		navIndex = 0
	} else if(window.location.pathname.indexOf('about') > -1){
		navIndex = 2
	} else if(window.location.pathname.indexOf('tools') > -1){
		navIndex = 3
	} else if(window.location.pathname.indexOf('nfts') > -1 ||window.location.pathname.indexOf('detail') > -1){
		navIndex = 1
	}
	if(navIndex > -1){
		$('.header__nav').find('.header__nav-item').eq(navIndex).addClass('active')
	}

	const keywords = sessionStorage.getItem('keywords');
	if (keywords) {
		$('.casting--search_input').val(keywords);
	}
	//Sign in 按钮绑定连接钱包事件
	$('.linkPurse').on('click', onLinkPurse);

	// input 绑定搜索事件
	$('.casting--search_input').on('input', function() {
		if (!$(this).val()) {
			sessionStorage.removeItem('keywords');
		}
	})
	$('.casting--search_btn').on('click', function() {
		console.log($('.casting--search_input').val())
		const keywords = $('.casting--search_input').val();
		sessionStorage.setItem('keywords', keywords)
		window.location.href="/nfts"
	})

	// 语言选择事件
	$('.language__dropdown-menu').find('li').on('click',function(){
		var language =$(this).attr('name');
		console.log(language);
		getCookie("userLanguage", language, {
			expires: 30,
			path: '/'
		});
		location.reload();
	})

	//点击铸件按钮
	$('.header__action-btn--signin').on('click',function(){
		const accountsKey = getCookie('accountsKey');
		console.log(accountsKey)
		if(accountsKey){
			window.location.href="/create"
		} else {
			onLinkPurse('create')
		}

	})

	const avatar = getCookie('avatar');

	if(avatar){
		$('.nav_avatar').remove();
		$('.linkPurse').css({
			'height':'40px',
			'width': '40px',
			'border-radius':'100%',
			'background': 'url('+avatar+') center center / cover no-repeat'
		})

	}


	$('.addNetwork').on('click',function(){
		addNetwork()
	})



	checkPurseStatus();
})

async function addNetwork(){
	try {
		await window.ethereum.request({
			method: 'wallet_switchEthereumChain',
			params: [{
				chainId: '0x5532c'
			}],
		});
	} catch (switchError) {
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
					}]
				});
			} catch (addError) {
				// handle "add" error
			}
		}
	}
}

//连接钱包
async function onLinkPurse(type) {

	if (metamask && typeof metamask.isInstall == 'function' && metamask.isInstall()) {
		console.log(metamask.isInstall())
		try {

			const accounts = await window.ethereum.request({
				method: "eth_requestAccounts",
			});
			const res = await userProfile({
				wallet: accounts[0]
			});
			$('.nav_avatar').attr('src',res.avatar)
			getCookie('avatar',res.avatar)
			getCookie('accountsKey', accounts[0]);

			if(type == 'create'){
				window.location.href="/create"
			} else if(type == 'reload'){
				window.location.reload()
			} else {
				window.location.href="/profile"
			}
		} catch (error) {
			$('.nav_avatar').attr('src',"https://cdn.weipaitang.com/static/2021092848a83b7c-aa1e-3b7caa1e-6ed2-c95c8a7a4e82-W650H650")
			getCookie('token','');
			getCookie('avatar','')
			getCookie('accountsKey', '');
			console.error(error);

		}
	} else {
		const currentUrl = new URL(window.location.href);
		const forwarderOrigin = currentUrl.hostname === "localhost" ? "http://localhost:9010" : undefined;
		const onboarding = new MetaMaskOnboarding({
			forwarderOrigin
		});
		onboarding.startOnboarding();
	}
}

//页面打开判断是否连接钱包
function checkPurseStatus() {
	if (metamask && typeof metamask.isInstall == 'function' && metamask.isInstall()) {
		window.ethereum.on('chainChanged', chainChanged);
		window.ethereum.on('accountsChanged', accountsChanged);

		setTimeout(() => {
			const chainId = window.ethereum.chainId;
			getCookie('chainId',chainId)

			const key = window.ethereum.selectedAddress;

			const accountsKey = getCookie('accountsKey');
			if (key && key == accountsKey) {
				getCookie('accountsKey',accountsKey)
			} else {
				// 再次检查是否连接了钱包
				setTimeout(() => {
					const selectedAddress = window.ethereum.selectedAddress;
					if (selectedAddress == accountsKey) {
						getCookie('accountsKey',accountsKey)
					} else {
						console.log(accountsKey)
						getCookie('accountsKey','');
						getCookie('token','');
						getCookie('avatar','')

					}
				}, 1000);
			}

		}, 1000);
	}
}

function chainChanged(chainId) {
	//0x1主网
	console.log(chainId);
	getCookie('chainId',chainId)
	if ((chainId == '0x1' && location.href.split("://")[0] === "http") || (chainId != '0x1' && location.href.split(
		"://")[0] === "https")) {
		showModal({
			message: $.i18n.prop('toast-DoesNotSupport'),
			type: 'error'
		})

	} else {
		window.location.reload();
	}
}

async function accountsChanged(accounts) {
	console.log(accounts)
	getCookie('avatar','')
	getCookie('token', '')
	if (accounts.length === 0) {
		getCookie("accountsKey", '')
		window.location.href = '/'
	} else {
		getCookie("accountsKey", accounts[0])
		await createUser({
			wallet: accounts[0],
			referrerWallet: getCookie('referrerWallet')
		})

		const res = await userProfile({
			wallet: accounts[0]
		})

		$('.nav_avatar').attr('src', res.avatar)

		getCookie('avatar',res.avatar);


		window.location.reload();
	}
}
