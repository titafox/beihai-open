var vm = new Vue({
	el: '#content',
	data: {
		userinfo: {},
		// view: 'casting-type',
		viewlist: ['casting-type', 'casting-single', 'collection-cover', 'casting-import', 'casting-multiple'],
		step: 0,
		mintDenied: false
	},
	created() {
		const accountsKey = getCookie('accountsKey');
		if(!accountsKey){
			onLinkPurse('reload')
		}else {
			this.getPression();
		}
		
	},
	computed: {
		view() {
			return this.viewlist[this.step]
		}
	},
	methods: {
		getPression() {
			isUpload({
				authType: 1
			}).then(res => {
				if (res.bol) {
					// if (type == "import") {
						//这里直接显示导入页面
					// 	this.step = 3;
					// }
					this.getProfileInfo();
					this.mintDenied = false
				} else {
					this.mintDenied = true
				}
			})
		},
		async getProfileInfo() {
			let res = {}
			const accountsKey = getCookie('accountsKey');
			res = await userCenter({
				wallet: accountsKey
			})
			
			if(res.userName == res.wallet){
				res.userName = formatWallet(res.userName)
			}
			this.userinfo = res;
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
		changeView() {
			console.log(arguments[0])
			this.step = arguments[0];
		}
	}
})
