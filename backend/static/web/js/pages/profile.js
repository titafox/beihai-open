var vm = new Vue({
	el: '#content',
	data: {
		isOwn: false, //是否是用户本人的个人页面
		wallet: '',
		userinfo: {},
		navlist: [{
				name: 'profile-own',
				cardlist: [],
				isEnd: false
			},
			{
				name: "profile-created",
				cardlist: [],
				isEnd: false
			}
		],
		cardlist: [],
		authorPage: 1,
		ownerPage: 1,
		isName: false,
		isDesc: false,
		isEdit: false,
		currentNav: 0,
		step: 0,
		viewlist: ['casting-single', 'collection-cover', 'casting-multiple'],
		collectionUri: '',
		uri:''
	},
	computed: {
		view() {
			return this.viewlist[this.step]
		},
		desclength() {
			const len = this.userinfo.signature?.length;
			return len;
		},
		descdisabled() {
			if(this.userinfo.signature?.length > 500){
				this.userinfo.signature = this.userinfo.signature.slice(0,500)
				return true
			}
			return false
		},
		inviteLink(){
			const link = `${window.location.origin}/index.html?referrerWallet=${this.userinfo.wallet}`
			return link
		}
	},
	created() {
		const accountsKey = getCookie('accountsKey');
		if(!accountsKey){
			onLinkPurse('reload')
		} else {
			const id = getQueryString('id');
			if (!id) {
				this.isOwn = true
			} else {
				this.wallet = id
			}

			this.getProfileInfo();
		}

	},

	methods: {
		async getProfileInfo() {
			let res = {}
			if (this.isOwn) {
				const accountsKey = getCookie('accountsKey');
				res = await userCenter({
					wallet: accountsKey
				})
			} else {
				const wallet = this.wallet;
				res = await userProfile({
					wallet: wallet
				})
			}

			console.log(res.userName)

			if(res.userName == res.wallet){
				res.userName = formatWallet(res.userName)
			}

			this.userinfo = res;
			this.getWorksOwner();
			this.getAuthorCasting()
		},
		copyWallet() {
			const input = document.createElement('input');
			input.value = this.userinfo.wallet;
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
		logout() {
			getCookie('token', '');
			getCookie('accountsKey', '');
			getCookie('avatar', '')
			window.location.href = "index.tmpl"
		},
		changeNavItem(i) {
			console.log(i)
			this.isEdit = false;
			this.currentNav = i;
			if (i == 1) {
				this.authorPage = 1;
				this.getAuthorCasting()
			}
		},
		changeView(step){
			this.step = step
		},
		changeEditStatus(){
			this.isEdit = !this.isEdit;
			if(!arguments[0]){
				const currentNav = this.currentNav;
				console.log(currentNav)
				if (currentNav == 1) {
					this.authorPage = 1;
					this.getAuthorCasting()
				} else {
					this.ownerPage = 1;
					this.getWorksOwner();
				}
			}

		},
		async getWorksOwner() {
			console.log('owner')
			const res = this.isOwn ? await worksOwner({
				page: this.ownerPage,
			}) : await otherWorksOwner({
				page: this.ownerPage,
				owner: this.wallet
			})

			if (this.ownerPage == 1) {
				this.navlist[0].cardlist = res.list;
			} else {
				this.navlist[0].cardlist = this.navlist[0].cardlist.concat(res.list);
			}
			this.ownerPage = res.page;
			this.navlist[0].isEnd = res.isEnd;

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
		},
		async getAuthorCasting() {
			console.log('create')
			const res = this.isOwn ? await authorCasting({
				page: this.authorPage,
			}) : await otherAuthorCasting({
				page: this.authorPage,
				owner: this.wallet
			})
			if (this.authorPage == 1) {
				this.navlist[1].cardlist = res.list;
			} else {
				this.navlist[1].cardlist = this.navlist[1].cardlist.concat(res.list);
			}
			this.authorPage = res.page;
			this.navlist[1].isEnd = res.isEnd;

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
		},

		getCardList() {
			let currentNav = this.currentNav
			console.log(currentNav)
			if (currentNav == 0) {
				this.getWorksOwner();
			} else {
				this.getAuthorCasting();
			}
		},
		cardClick(i, j) {
			const cardlist = this.navlist[i].cardlist;
			const isOwn = this.isOwn;
			const cardinfo = cardlist[j]
			const uri = cardinfo.uri;
			const collectionUri = cardinfo.collectionUri;
			if (collectionUri) {
				this.collectionUri = collectionUri;
				this.authorPage = 1
				this.getCollectionWorks()
			} else {
				//点击单个项目
				window.location.href = `detail.html?id=${uri}`
			}


			// if (isOwn) {
			// 	//查看本人的个人中心

			// 	const collectionUri = cardinfo.collectionUri;
			// 	if (collectionUri) {
			// 		this.collectionUri = collectionUri;
			// 		this.authorPage = 1
			// 		this.getCollectionWorks()
			// 	} else {
			// 		//点击单个项目
			// 		window.location.href = `detail.tmpl?id=${uri}`
			// 	}

			// } else {


			// 	//查看他人的个人中心
			// 	window.location.href = `detail.tmpl?id=${uri}`
			// }
		},
		addCasting(i, j) {


			const cardlist = this.navlist[i].cardlist;
			const collectionUri = cardlist[j].collectionUri;
			this.collectionUri = collectionUri;
			this.isEdit = true;
			this.step = 2;
		},
		editCasting(i, j) {
			const cardlist = this.navlist[i].cardlist;
			const cardinfo = cardlist[j]
			const uri = cardinfo.uri;
			const collectionUri = cardinfo.collectionUri;
			console.log(collectionUri)
			if (collectionUri) {
				this.isEdit = true;
				this.step = 1;
				this.collectionUri = collectionUri;
			} else {
				this.isEdit = true;
				this.step = 0;
				this.uri = uri
			}
		},
		getCollectionWorks() {
			collectionWorks({
				collectionUri: this.collectionUri,
				page: this.authorPage,
				owner: this.wallet
			}).then(res => {
				if (this.authorPage == 1) {
					this.navlist[1].cardlist = res.list;
				} else {
					this.navlist[1].cardlist = this.navlist[1].cardlist.concat(res.list);
				}
				this.authorPage = res.page;
				this.navlist[1].isEnd = res.isEnd;
			})
		},
		editName(){
			this.editInfo()
			this.isName = false;
		},
		editDesc(){
			this.editInfo()
			this.isDesc = false;
		},
		async editInfo(){

			const accountsKey = getCookie('accountsKey')
			await setUserProfile({
				wallet: accountsKey,
				userName: this.userinfo.name,
				signature: this.userinfo.signature,
				avatar: this.userinfo.avatar,
				backgroundImage: this.userinfo.background_image
			})
			getCookie('avatar',this.userinfo.avatar)
			$('.nav_avatar').attr('src',this.userinfo.avatar)
		},
		avatarClick(){
			if(!this.isOwn){
				return;
			}
			const uploadImg = new Upload({
				limit: 1, // 单次上传数量限制，默认9
				maxSize: 50 * 1024 * 1024, // 大小限制，公众号环境上传(env传wechat时)没有，默认10M
			});
			uploadImg.upload('', () => {
				//showloading
			}).then(res => {
				console.log(res.serverIds[0])
				// hideloading
				this.userinfo.avatar = res.serverIds[0];
				this.editInfo();
			}).catch((res) => {
				console.log(res)

			})
		},
		uploadBg(){
			const uploadImg = new Upload({
				limit: 1, // 单次上传数量限制，默认9
				maxSize: 50 * 1024 * 1024, // 大小限制，公众号环境上传(env传wechat时)没有，默认10M
				multiple: true
			});
			uploadImg.upload('', () => {
				//showloading
			}).then(res => {
				console.log(res.serverIds[0])
				// hideloading
				this.userinfo.background_image = res.serverIds[0];
				this.editInfo();
			}).catch((res) => {
				console.log(res)

			})
		},
		invite(){
			$('.invite-modal').removeClass('fadeIn').addClass('fadeOut');
			const input = document.createElement('input');
			console.log(window.location.origin)
			const val = `${window.location.origin}/index.html?referrerWallet=${this.userinfo.wallet}`
			input.value = val;
			input.style = 'position: fixed';
			document.body.appendChild(input);
			input.select();
			document.execCommand("copy");
			document.body.removeChild(input);
			showToast({
				type: 'success',
				duration: 5000,
				message: $.i18n.prop('toast-copySuccesseful')
			});
		},
	}
})
