var vm = new Vue({
	el: '#content',
	data: {
		cardlist: [],
		slideImg: [],
	},
	created() {
		let referrerWallet = getQueryString('referrerWallet');
		if (referrerWallet) {
			getCookie('referrerWallet', referrerWallet)
			bindUser({
				referrerWallet: referrerWallet
			}).then(res => {
				getCookie('referrerWallet', '')

			})
		}



		this.getRecommendAuthor()
		this.getRecommendCarousel()
	},
	methods: {
		getRecommendAuthor() {
			recommendAuthor().then(res => {
				for (let i = 0; i < res.length; i++) {
					let items = res[i].items;
					for (let j = 0; j < items.length; j++) {
						items[j].owner = {}
						items[j].owner.userName = res[i].user_name;
						items[j].owner.avatar = res[i].avatar;
						items[j].owner.wallet = res[i].wallet;
					}
					res[i].items = items;
				}
				this.cardlist = res
			}).catch(res => {
				console.log(res)
			});
		},
		getRecommendCarousel() {
			recommendCarousel().then(res => {
				this.slideImg = res

				
				
				setTimeout(() => {
					$('.index-carousel').owlCarousel({
						mouseDrag: true,
						touchDrag: true,
						dots: true,
						loop: true,
						autoplay: true,
						autoplayHoverPause: true,
						smartSpeed: 1500,
						autoHeight: true,
						items: 1,
						responsive: {
							0 : {
								margin: 20,
							},
							576 : {
								margin: 20,
							},
							768 : {
								margin: 30,
							},
							1200 : {
								margin: 30,
							},
						}
					});
				}, 20)
			}).catch(res => {
				console.log(res)
			});
		},
		heartClick(i, j) {
			console.log(i, j)
			let cardInfo = this.cardlist[i].items[j]
			console.log(cardInfo)
			const worksUri = cardInfo.uri;
			const status = cardInfo.isLike ? 1 : 0;
			castingLike({
				worksUri: worksUri,
				status: status
			}).then(res => {
				cardInfo.isLike = !cardInfo.isLike;
				cardInfo.like_num = cardInfo.isLike ? cardInfo.like_num + 1 : cardInfo.like_num - 1;
				this.cardlist[i].items.splice(j, 1, cardInfo)
			}).catch(res => {
				console.log(res)
			})
		},
		toProfile(wallet) {
			window.location.href = 'profile.html?id=' + wallet;
		}
	}
})
