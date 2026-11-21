var vm = new Vue({
	el: '#content',
	data: {
		uri: '',
		goodsDetail: {

		},
		carouselList: [],
		bill: [],
		chainId: ''
	},
	created() {
		const href = window.location.href;

		let referrerWallet = ''
		if(href.indexOf('/referrerWallet/') > -1){
			let referrerWalletArr = href.split('/referrerWallet/')
			referrerWallet = referrerWalletArr[1];


			this.uri = referrerWalletArr[0].split('/share/')[1]

		} else {
			const hrefArr = href.split('/detail/');
			this.uri =hrefArr[hrefArr.length-1];
		}




		this.getCastingDetails();

		this.getPriceInfo();

		if(referrerWallet){
			getCookie('referrerWallet', referrerWallet)
			bindUser({
				referrerWallet: referrerWallet
			}).then(res=>{
				getCookie('referrerWallet', '')

			})
		}
	},
	computed: {
		twitterLink() {
			// referrerWallet
			let link = '';
			if (Object.keys(this.goodsDetail).length == 0) {
				return link
			} else {
				let url = `${window.location.origin}/share/${this.goodsDetail.uri}`;
				let title = this.goodsDetail.title;
				title = title.split('#').join('')

				const referrerWallet = getCookie('accountsKey');
				if(referrerWallet){
					url = `${url}/referrerWallet/${referrerWallet}`
				}


				link = `https://twitter.com/intent/tweet?hashtags=NFT,Coltstail&text=${title}&url=${url}`;
				return link;
			}
		},
		facebookLink() {
			let link = '';
			if (Object.keys(this.goodsDetail).length == 0) {
				return link
			} else {

				let url = `${window.location.origin}/share/${this.goodsDetail.uri}`;
				const referrerWallet = getCookie('accountsKey');
				if(referrerWallet){
					url = `${url}/referrerWallet/${referrerWallet}`
				}

				link =
					`https://www.facebook.com/dialog/share?hashtags=NFT,Coltstail&app_id=1017701618797371&display=popup&href=${url}`

				return link;
			}

		},
		availableForPurchase() {
			let key = this.goodsDetail.owner?.wallet;
			let accountsKey = getCookie('accountsKey');
			const chainId = getCookie('chainId');
			this.chainId = chainId
			let typ = this.goodsDetail.contract_type;
			// if (!accountsKey || key == accountsKey || this.goodsDetail.goodsDetail === 1 || (chainId == '0x1' &&location.href.split("://")[0] === "http") || (chainId != '0x1' && location.href.split("://")[0] === "https")) {
			// 	return false
			// }


			if (!accountsKey || key == accountsKey || this.goodsDetail.goodsStatus === 1||(typ=='1' && chainId !='0x5532c') || (typ=='0' && chainId == '0x1' &&location.href.split("://")[0] === "http") || (typ=='0' && chainId != '0x1' && location.href.split("://")[0] === "https")) {
				return false
			}



			return true
		}
	},
	methods: {
		async getCastingDetails() {
			const uri = this.uri;
			const res = await worksDetail({
				uri: uri
			})
			this.goodsDetail = res;



			this.getRecommendList();
		},
		getRecommendList() {
			const category = this.goodsDetail.category;
			const collection_uri = this.goodsDetail.collection_uri;
			worksRecommend({
				page: 0,
				collection_uri: collection_uri,
				category: category
			}).then(res=>{
				console.log(res.list[0]);
				this.carouselList = res.list;
				setTimeout(()=>{
					console.log('render')
					console.log($('.main__carousel').length)
					$('.main__carousel').owlCarousel({
						mouseDrag: true,
						touchDrag: true,
						dots: true,
						loop: true,
						autoplay: true,
						autoplayHoverPause: true,
						autoplayTimeout: 5000,
						smartSpeed: 600,
						margin: 20,
						autoHeight: true,
						responsive: {
							0 : {
								items: 1,
							},
							576 : {
								items: 2,
							},
							768 : {
								items: 3,
								margin: 30,
							},
							992 : {
								items: 4,
								margin: 30,
							},
							1200 : {
								items: 4,
								margin: 30,
								mouseDrag: false,
								dots: false,
							},
						},
					});


				},20)


			})


		},
		getPriceInfo() {
			const uri = this.uri;
			castingTransaction({
				uri: uri
			}).then(res => {
				let bill = res.bill;
				for (let i = 0; i < bill.length; i++) {
					bill[i].createdTime = time(bill[i].createdTime);
					bill[i].buyerWallet = formatWallet(bill[i].buyerWallet)
				}

				this.bill = bill;


			})
		},
		heartClick() {
			let worksUri = '';
			let status = null;
			let uri  = arguments[0]
			let index = arguments[1]
			console.log(index)
			if(index > -1){
				worksUri = uri;
				var cardInfo = this.carouselList[index]
				status = cardInfo.isLike ? 1 : 0;
			} else {

				worksUri = this.uri;
				status = this.goodsDetail.isLike ? 1 : 0;
			}
			castingLike({
				worksUri: worksUri,
				status: status
			}).then(res => {
				if(index > -1){
					console.log(cardInfo)
					cardInfo.isLike = !cardInfo.isLike;
					cardInfo.like_num = cardInfo.isLike ? cardInfo.like_num + 1 : cardInfo.like_num - 1;
					this.carouselList.splice(index,1,cardInfo)
				} else {
					this.goodsDetail.isLike = !this.goodsDetail.isLike;
					this.goodsDetail.likeNum = this.goodsDetail.isLike ? this.goodsDetail.likeNum + 1 : this
						.goodsDetail.likeNum - 1;
				}

			}).catch(res => {
				console.log(res)
			})
		},
		copyLink() {
			const input = document.createElement('input');
			let url = `${window.location.origin}/share/${this.goodsDetail.uri}`;
			const referrerWallet = getCookie('accountsKey');
			if(referrerWallet){
				url = `${url}/referrerWallet/${referrerWallet}`
			}

			input.value = url;

			input.style = 'position: fixed';
			document.body.appendChild(input);
			input.select();
			document.execCommand("copy");
			document.body.removeChild(input);
			showToast({
				type: 'success',
				duration: 5000,
				message: $.i18n.prop('toast-copySuccesseful')
			})
		},
		refresh() {
			window.location.reload()
		},
		async toBuy() {
			console.log('click')

			const canBuy = this.availableForPurchase;
			if (!canBuy) {
				return;
			}

			if (!metamask?.isInstall) {
				showModal({
					type: 'error',
					message: $.i18n.prop('toast-pleaseInstallThePlugin'),
				});
			}

			let typ = this.goodsDetail.contract_type;




			if(typ == '0'){
				this.toBuyNft()
			} else {
				// 调用合约里的balanceOf查询余额
				this.toBuyDogTicket()
			}



		},
		async toBuyNft(){
			let wallet = getCookie('accountsKey')

			if(!wallet){
				onLinkPurse()
				return;
			}
			const balance = await metamask?.getBalance(wallet);
			console.log(balance)
			if (balance * 1000000000000000000 < this.goodsDetail.price) {
				showModal({
					type: 'error',
					message: $.i18n.prop('toast-balance'),
				});
				return
			}

			showToast({
				type: 'loading',
				title: 'loading...',
				message: $.i18n.prop('toast-info'),
			})



			let buyData = await tradeBuy({
				worksUri: this.goodsDetail.uri,
				buyer: wallet
			})

			buyData.percentage_basis_points = buyData.percentage_basis_points.toString();
			buyData.token_uri = buyData.metaData.toString();
			const sellOrder = await metamask.sellOrderNft({
				...buyData
			});
			tradeBuyBefore({
				worksUri: this.goodsDetail.uri,
				hash: sellOrder.hash
			})

			await metamask.getStatusWithHash(sellOrder.hash, () => {
				hideToast()
				window.location.href="/profile"
			});
		},
		async toBuyDogTicket(){
			let wallet = getCookie('accountsKey');
			if(!wallet){
				onLinkPurse()
				return;
			}

			const balance = await metamask?.getBalanceOfCTS(wallet);

			console.log(balance)
			let price = ethers.utils.formatEther(this.goodsDetail.price);
			// let value = this.goodsDetail.price / 1000000000000000000
			if (balance < price) {
				showModal({
					type: 'error',
					message: $.i18n.prop('toast-balance'),
				});
				return
			}

			showToast({
				type: 'loading',
				title: 'loading...',
				message: $.i18n.prop('toast-info'),
			})


			console.log(this.goodsDetail.price);
			// let price = ethers.utils.formatEther(value);

			const data = await metamask?.ticketApprove({
				address: wallet,
				price: price//this.goodsDetail.price
			})
			console.log(data)
			console.log(data.hash)

			await metamask.getStatusWithHash(data.hash, () => {
				this.afterSetApprove()
			});







		},
		async afterSetApprove(){
			const tokeId = this.goodsDetail.TokenInfo.tokenId;
			const data = await metamask?.ticketBuy({
				tokenId:tokeId
			})


			tradeBuyBefore({
				worksUri: this.goodsDetail.uri,
				hash: data.hash
			})

			await metamask.getStatusWithHash(data.hash, () => {
				hideToast()
				window.location.href="/profile"
			});
		}
	}
})
