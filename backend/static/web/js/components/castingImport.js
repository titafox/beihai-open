//导入铸件
Vue.component('casting-import', {
	template: `
		<div class="casting-import">
			<div class="tab-content">
				<div class="tab-pane fade show active" role="tabpanel">
					<div class="row row--grid">
						<div class="col-12">
							<form action="#" class="sign__form sign__form--profile">
								<div class="row">
									<div style="width:100%" v-show="!isDetail">
										<div class="col-12">
											<h4 class="sign__title i18n" name="create-creationcode">
												{{$.i18n.prop('create-creationcode')}}
											</h4>
										</div>
										<div class="col-12">
											<div class="sign__group sign__group-container">
												<input v-model="code" class="sign__input">
											</div>
										</div>
										
										<div class="col-12">
											<div class="row">
												<div class="col-3">
													<button class="sign__btn i18n" name="create-submit" type="button" @click="codeHandler">
														{{$.i18n.prop('create-submit')}}
													</button>
												</div>
											</div>
										</div>
									</div>
									
								</div>
								
								<div v-show="isDetail" style="width:100%">
								
									<div class="col-12">
											<h4 class="sign__title i18n" name="create-uploadFiles">
												{{$.i18n.prop('create-uploadFiles')}}
											</h4>
										</div>
										<div class="col-12">
											<div class="sign__group">
												<div class="upload__img-container">
													<div class="upload__img">
														<div class="exit_image" :style="{background:'url('+imgUrl+') center center / contain no-repeat'}"></div>
													</div>
												</div>
											</div>
										</div>
										<div class="col-12">
											<div class="sign__group">
												<label class="sign__label i18n" name="create-title">
													{{$.i18n.prop('create-title')}}
												</label>
												<div class="sign__input-container">
													<input disabled="disabled" v-model="title" :placeholder="$.i18n.prop('create-title')" type="text" class="sign__input i18n-input" select-name="create-pleaseEnterName" select-attr="placeholder">
													<img src="img/components/lock.png"/>
												</div>
											</div>
										</div>
										
										<div class="col-12">
											<div class="sign__group">
												<label class="sign__label i18n" name="create-desc">
													{{$.i18n.prop('create-desc')}}
												</label>
												<div class="sign__input-container" style="position:relative">
													<textarea disabled="disabled" v-model="desc" name="text" :placeholder="$.i18n.prop('create-pleaseEnterDesc')" class="sign__textarea i18n-input" select-name="create-pleaseEnterDesc" select-attr="placeholder"></textarea>
													<img src="img/components/lock.png"/>
													<span style="position: absolute;color: #FFFFFF;right: 20px;bottom: 0;">{{desclength}} / 500</span>
												</div>
											</div>
										</div>
										<div class="col-12">
											<div class="sign__group sign__group-container">
												<label class="sign__label i18n" name="common-price">
													{{$.i18n.prop('common-price')}}
												</label>
												<div class="sign__input-container">
													<input disabled="disabled" v-model="price" :placeholder="$.i18n.prop('create-pleaseEnterPrice')" maxlength="8" type="number" class="sign__input i18n-input" select-name="create-pleaseEnterPrice" select-attr="placeholder">
													<span>ETH</span>
													<img src="img/components/lock.png"/>
												</div>
											</div>
										</div>
										
										<div class="col-12">
											<div class="sign__group">
												<label class="sign__label i18n" name="create-classification" >
													{{$.i18n.prop('create-classification')}}
												</label>
												<div class="sign__input-container">
													<select name="select" class="sign__select" v-model="category" disabled="disabled">
														<option v-for="(item,i) in categorylist" :key="item.id" :value="item.id" >{{item.categoryName}}</option>
													</select>
													<img src="img/components/lock.png"/>
												</div>
												
											</div>
										</div>
										
										<div class="col-12">
											<div class="sign__group">
												<label class="sign__label i18n" name="create-remuneration">
													{{$.i18n.prop('create-remuneration')}}
												</label>
												<div  class="sign__input-container">
													<select name="select" class="sign__select" v-model="royaltRate" disabled="disabled">
														<option value="4">4%</option>
														<option value="5">5%</option>
														<option value="6">6%</option>
														<option value="7">7%</option>
														<option value="8">8%</option>
														<option value="9">9%</option>
														<option value="10">10%</option>
													</select>
													<img src="img/components/lock.png"/>
												</div>
											</div>
										</div>
										
										
										
										
										
										<div class="col-12">
											<div class="row">
												<div class="col-3">
													<button class="sign__btn i18n" name="create-create" type="button" @click="mintHandler">
														{{$.i18n.prop('create-save')}}
													</button>
												</div>
											</div>
										</div>
									
								</div>
							</form>
						</div>
					</div>
				</div>
			</div>
		</div>
	`,
	data: function() {
		return {
			code: '',
			imgUrl: '',
			title: '',
			desc: '',
			price: '',
			royaltRate: 4,
			category: '',
			isloading: false,
			categorylist: [],
			signInfo: '',
			isDetail: false
		}
	},
	created() {
		this.getCategoryList()
	},
	watch: {
		signInfo(newval, oldval) {
			if (newval != oldval && newval) {
				this.submitCode()
			}
		}
	},
	computed:{
			desclength(){
				const len = this.desc.length;
				return len;
			}
	},
	methods: {
		getCategoryList() {
			categoryList().then(res => {
				this.categorylist = res.list
				if (res.list.length > 0) {
					this.category = res.list[0].id;
				}
			}).catch(res => {
				console.log(res.list)
			})
		},
		codeHandler() {
			const code = this.code;
			console.log(this.code)
			if (!code) {
				return
			}
			importPre({
				code: code
			}).then(res => {
				this.imgUrl = res.cover;
				this.title = res.title;
				this.desc = res.description;
				this.royaltRate = res.royaltRate;
				this.category = res.category;
				this.price = res.price;
				this.isDetail = true
			})
		},
		async mintHandler() {
			const accountsKey = getCookie('accountsKey')
			const balance = await metamask.getBalance(accountsKey);
			console.log(balance)
			if (balance * 1000000000000000000 == 0) {
				showModal({
					type: 'error',
					message: $.i18n.prop('toast-balance'),
					duration: 1500
				});
				return
			}

			const code = this.code;
			if (!code) {
				return
			}

			try {
				await this.authorization();
			} catch (err) {
				console.log(err)
			}
		},
		async authorization() {
			try {
				showToast({
					type: 'loading',
					message: $.i18n.prop('create-nftStatus'),
				})
				const isChained = await metamask?.hasChained();
				if (!isChained) {
					const str = await metamask?.createSignMessage();
					const data = await metamask?.getHashData();
					await metamask?.getStatusWithHash(data.hash, () => {
						this.signInfo = str;
					});
				} else {
					// 签名
					hideToast();
					setTimeout(async () => {
						try {
							showToast({
								type: 'loading',
								message: $.i18n.prop('create-getSign'),
							})
							const str = await metamask?.createSignMessage();
							this.signInfo = str;
						} catch (err) {
							showModal({
								type: 'error',
								message: err.message,
								duration: 1500
							});
							this.signInfo = '';
						}
					}, 800)
				}
			} catch (err) {
				//TODO err 设置
				showModal({
					type: 'error',
					message: $.i18n.prop('toast-deductionFailed')
				});
				this.signInfo = '';
			}
		},
		submitCode() {
			const code = this.code;
			importCollection({
				code: code
			}).then(() => {
				showToast({
					message: $.i18n.prop('create-createdSuccessfully'),
					type: 'success',
					duration: 5000,
					callback: ()=>{
						window.location.href = 'profile.html'
					}
				});
			
			})
			this.signInfo = ''
		}
	}
})
