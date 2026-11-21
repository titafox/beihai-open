var vm = new Vue({
	el: '#content',
	data: {
		navList: [],
		page: 1,
		category: '',
		keywords: '',
		isEnd:true,
		total: '',
		cardlist:[]
	},
	created() {
		setTimeout(()=>{
			this.getCategoryList();
			this.keywords = $('.casting--search_input').val()
		},100)
	},
	methods: {
		getCategoryList() {
			categoryList().then(res => {
				this.navList = [];
				let list = [];
				for (let i = 0; i < res.list.length; i++) {
					if (res.list[i].itemCount > 0) {
						list.push(res.list[i])
					}
				}
				list.unshift({
					id: '',
					categoryName: $.i18n.prop('nfts-all')
				})
				this.navList = list;

				setTimeout(() => {
					$('.card__cover--carousel').owlCarousel({
						mouseDrag: true,
						touchDrag: true,
						dots: true,
						loop: true,
						autoplay: true,
						autoplayHoverPause: true,
						autoplayTimeout: 5000,
						autoplaySpeed: 800,
						smartSpeed: 800,
						margin: 20,
						items: 1,
					});
				}, 20)
				this.getCardList();
				
			}).catch(res => {
				console.log(res)
			})
		},
		changeNavItem(id){
			this.category = id;
			this.page = 1;
			this.getCardList()
		},
		async getCardList() {
			const page = this.page;
			console.log(page)
			const category = this.category;
			const keywords = this.keywords;
			const res = await castingList({
				page: page,
				title: keywords,
				category: category
			})
			
			console.log(res)

			if (page == 1) {
				this.cardlist = res.list;
			} else {
				this.cardlist = this.cardlist.concat(res.list);
			}

			this.isEnd = res.isEnd;
			this.page = res.page;
			this.total = res.total;
		},
		heartClick() {
			let worksUri = '';
			let status = null;
			let uri  = arguments[0]
			if(uri){
				worksUri = uri;
				var index = arguments[1];
				var cardInfo = this.cardlist[index]
				status = cardInfo.isLike ? 1 : 0;
			} else {
				worksUri = this.uri;
				status = this.goodsDetail.isLike ? 1 : 0;
			}
			castingLike({
				worksUri: worksUri,
				status: status
			}).then(res => {
				if(uri){
					cardInfo.isLike = !cardInfo.isLike;
					cardInfo.like_num = cardInfo.isLike ? cardInfo.like_num + 1 : cardInfo.like_num - 1;
					this.cardlist.splice(index,1,cardInfo)
				} else {
					this.goodsDetail.isLike = !this.goodsDetail.isLike;
					this.goodsDetail.likeNum = this.goodsDetail.isLike ? this.goodsDetail.likeNum + 1 : this
						.goodsDetail.likeNum - 1;
				}
				
			}).catch(res => {
				console.log(res)
			})
		},
		linkToDetail(id){
			window.location.href=`/detail/${id}`
		}
	}
})
