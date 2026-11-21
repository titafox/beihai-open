//单个铸件

Vue.component('casting-single', {
	template: `
		<div class="casting-single">
			<div class="tab-content">
				<div class="tab-pane fade show active" role="tabpanel">
					<div class="row row--grid">
						<div class="col-12">
							<form action="#" class="sign__form sign__form--profile">
								<div class="row">
									<div class="col-12">
										<h4 class="sign__title i18n" name="create-uploadFiles">
											{{$.i18n.prop('create-uploadFiles')}}
										</h4>
									</div>
									<div class="col-12">
										<div class="sign__group">
											<div class="upload__img-container" :class="(uri)?'no_upload':''" @click="uploadCastingImg">
												<div class="upload__img">
													<template v-if="!imgUrl">
														<div v-if="!isloading" class="no_image" >
															<img src="img/components/upload-icon.svg"/>
															<span class="i18n" name="create-uploadImage">
																{{$.i18n.prop('create-uploadImage')}}
															</span>
														</div>
														<div v-else class="loading">
															<img src="img/common/loading.svg"/>
														</div>
													</template>
													<template v-else>
														<div class="exit_image" :style="{background:'url('+imgUrl+'/640) center center / contain no-repeat'}"></div>
													</template>
												</div>
											</div>
										</div>
									</div>
									<div class="col-12">
										<h4 class="sign__title i18n" name="create-workDetails">
											{{$.i18n.prop('create-workDetails')}}
										</h4>
									</div>
									<div class="col-12">
										<div class="sign__group">
											<label class="sign__label i18n" name="create-title">
												{{$.i18n.prop('create-title')}}
											</label>
											
											<input v-if="!uri" v-model="title" :placeholder="$.i18n.prop('create-pleaseEnterName')" type="text" class="sign__input i18n-input" select-name="create-pleaseEnterName" select-attr="placeholder">
											
											
											<div v-if="uri" class="sign__input-container">
												<input disabled="disabled" v-model="title" :placeholder="$.i18n.prop('create-pleaseEnterName')" type="text" class="sign__input i18n-input" select-name="create-pleaseEnterName" select-attr="placeholder">
												<img src="img/components/lock.png"/>
											</div>
										
										</div>
									</div>
									
									<div class="col-12">
										<div class="sign__group" style="position:relative">
											<label class="sign__label i18n" name="create-desc">
												{{$.i18n.prop('create-desc')}}
											</label>
											<div v-if="goodsStatus == 2" class="sign__input-container">
												<textarea disabled="disabled" v-model="desc"  :placeholder="$.i18n.prop('create-pleaseEnterDesc')" name="text" class="sign__textarea i18n-input" select-name="create-pleaseEnterDesc" select-attr="placeholder"></textarea>
												<img src="img/components/lock.png"/>
											</div>
											
											<textarea v-else :descdisabled="descdisabled"  v-model="desc" :placeholder="$.i18n.prop('create-pleaseEnterDesc')" name="text" class="sign__textarea i18n-input" select-name="create-pleaseEnterDesc" select-attr="placeholder"></textarea>
											
											<span style="position: absolute;color: #FFFFFF;right: 20px;bottom: 0;">{{desclength}} / 500</span>
										
										</div>
									</div>
									
									<div class="col-12">
										<div class="sign__group sign__group-container">
											<label class="sign__label i18n" name="common-price">
												{{$.i18n.prop('common-price')}}
											</label>
											
											<div v-if="goodsStatus == 2" class="sign__input-container">
												<input disabled="disabled" v-model="price" maxlength="8" type="number" :placeholder="$.i18n.prop('create-pleaseEnterPrice')" class="sign__input i18n-input" select-name="create-pleaseEnterPrice" select-attr="placeholder">
												<span>ETH</span>
												<img src="img/components/lock.png"/>
											</div>
											<div v-else  class="sign__input-container">
												<input v-model="price" maxlength="8" type="number" :placeholder="$.i18n.prop('create-pleaseEnterPrice')" class="sign__input i18n-input" select-name="create-pleaseEnterPrice" select-attr="placeholder">
												<span>ETH</span>
											</div>
											
										</div>
									</div>
									
									<div class="col-12">
										<div class="sign__group">
											<label class="sign__label i18n" name="create-classification" >
												{{$.i18n.prop('create-classification')}}
											</label>
											<select v-if="!uri" name="select" class="sign__select" v-model="category">
												<option v-for="(item,i) in categorylist" :key="item.id" :value="item.id" >{{item.categoryName}}</option>
											</select>
											
											<div v-if="uri" class="sign__input-container">
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
											<select  v-if="!uri" name="select" class="sign__select" v-model="royaltRate">
												<option value="4">4%</option>
											    <option value="5">5%</option>
											    <option value="6">6%</option>
											    <option value="7">7%</option>
											    <option value="8">8%</option>
											    <option value="9">9%</option>
											    <option value="10">10%</option>
											</select>
											
											<div v-if="uri" class="sign__input-container">
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
									<div class="col-12" style="display: flex">
										<button v-if="!uri" class="sign__btn i18n" name="create-create" type="button" @click="mintHandler">
											{{$.i18n.prop('create-create')}}
										</button>
										<button v-if="uri && goodsStatus == 1" type="button" class="sign__btn" @click="updateCasting">
											{{$.i18n.prop('create-sell')}}
										</button>
										<button v-if="uri && goodsStatus == 2" type="button" class="sign__btn" @click="updateCasting">
											{{$.i18n.prop('create-soldOut')}}
										</button>
										<button style="margin-left:10px" v-if="uri && goodsStatus != 0" type="button" class="sign__btn" @click="$emit('viewclick','1')">
											{{$.i18n.prop('create-back')}}
										</button>
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
			imgUrl: '',
			title: '',
			desc: '',
			price: '',
			royaltRate: 4,
			category: '',
			isloading: false,
			categorylist: [],
			signInfo: '',
			goodsStatus: 0
		}
	},
	props: ['uri'],
	watch: {
		signInfo(newval, oldval) {
			if (newval != oldval && newval) {
				this.createCasting()
			}
		},
		uri(newval,oldval){
			console.log(newval,oldval)
			this.getCategoryList()
		}
	},
	computed:{
			desclength() {
				const len = this.desc.length;
				return len;
			},
			descdisabled() {
				if(this.desc.length > 500){
					this.desc = this.desc.slice(0,500)
					return true
				}
				return false
			}
	},
	created() {
		this.getCategoryList()
		console.log(this.uri)
	},
	methods: {
		getCategoryList() {
			categoryList().then(res => {
				this.categorylist = res.list
				if (res.list.length > 0) {
					this.category = res.list[0].id;
				}

				if (this.uri) {
					this.getCasting()
				}
			}).catch(res => {
				console.log(res.list)
			})
		},

		getCasting() {
			worksDetail({
				uri: this.uri
			}).then(res => {
				this.title = res.title
				this.desc = res.description
				this.category = res.category
				this.imgUrl = res.cover
				this.goodsStatus = res.goodsStatus
				this.royaltRate = res.royalt_rate
				this.price = (res.price / 1000000000000000000)
			})
		},
		uploadCastingImg() {
			console.log('uploadImg')
			if (this.uri) {
				return
			}
			const uploadImg = new Upload({
				limit: 1, // 单次上传数量限制，默认9
				maxSize: 50 * 1024 * 1024, // 大小限制，公众号环境上传(env传wechat时)没有，默认10M
			});
			uploadImg.upload('', () => {
				this.isloading = true
			}).then(res => {
				this.isloading = false
				this.imgUrl = res.serverIds[0];
			}).catch((res) => {
				console.log(res)
			})
		},
		//铸件
		async mintHandler() {
			if (!this.uri) {
				const accountsKey = getCookie('accountsKey')
				const balance = await metamask.getBalance(accountsKey);
				console.log(balance)
				if (balance * 1000000000000000000 == 0) {
					showModal({
						type: 'error',
						message: $.i18n.prop('toast-balance'),
						
					});
					return
				}

				const {
					imgUrl,
					title,
					royaltRate,
					price
				} = this;

				if (price < 0.01 || !price) {
					showModal({
						type: 'error',
						message: $.i18n.prop('toast-lowest'),
						duration: 1500
					})
					return
				}

				if ((price.toString().indexOf('.') > -1 && price.toString().length > 8) || price >=
					10000000) {
					showModal({
						type: 'error',
						message: $.i18n.prop('toast-maxlength'),
						duration: 1500
					})
					return
				}

				if (!imgUrl || !title || !royaltRate) {
					showModal({
						type: 'error',
						message: $.i18n.prop('toast-pleaseEnterCompleteInformation'),
						duration: 1500
					});
					return
				}

				try {
					await this.authorization();
				} catch (err) {
					console.log(err)
				}



			} else {
				this.updateCasting()
			}
		},
		//铸件
		createCasting() {
			const {
				imgUrl,
				title,
				royaltRate,
				desc,
				price,
				category,
				signInfo
			} = this;
			const accountsKey = getCookie('accountsKey');
			saveCasting({
				cover: imgUrl,
				title: title,
				description: desc,
				royaltRate: +royaltRate,
				price: +price,
				author: accountsKey,
				sign: signInfo,
				category: category
			}).then(() => {
				this.signInfo = '';
				console.log('success')
				showToast({
					message: $.i18n.prop('create-createdSuccessfully'),
					duration: 5000,
					type: 'success',
					callback: () => {
						window.location.href = "profile.html"
					}
				});
			});
		},
		//修改铸件
		updateCasting() {
			worksEdit({
				uri: this.uri,
				description: this.desc,
				price: this.price,
				status: this.goodsStatus === 1 ? 2 : 1,
				category: this.category
			}).then(() => {
				// TODO 修改
				showToast({
					message: $.i18n.prop('create-createdSuccessfully'),
					duration: 5000,
					callback: () => {
						this.$emit('viewclick')
					}
				})
			})
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
							
							});
							this.signInfo = '';
						}
					}, 800)
				}
			} catch (err) {
				//TODO err 设置
				showModal({
					type: 'error',
					message: $.i18n.prop('toast-deductionFailed'),
					duration: 1500
				});
				this.signInfo = '';
			}
		}
	}
})
