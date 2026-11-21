//多个铸件

Vue.component('casting-multiple', {
	template: `
		<div class="casting-multiple">
			<div class="tab-content">
				<div class="tab-pane fade show active" role="tabpanel">
					<div class="row row--grid">
						<div class="col-12">
							<form action="#" class="sign__form sign__form--profile">
								<div class="row">
									<div class="col-10" v-if="!multipleDenied">
										<div class="sign__group upload__box">
											<label class="sign__label i18n" name="create-selectUploadType">{{$.i18n.prop('create-selectUploadType')}}</label>
											<!--<select name="select" class="sign__select" v-model="uploadType">
												<option value="1" class="i18n" name="create-single">{{$.i18n.prop('create-single')}}</option>
												<option value="2" class="i18n" name="create-multiple">{{$.i18n.prop('create-multiple')}}</option>
											</select>-->
											
											
											 <div class="sign__input upload__input" disabled="true">
											              <span>{{uploadTypeShow}}</span>
											            </div>
											            <div class="upload__select">
											              <div class="upload__select_trangle">
											                <ul>
											                  <li @click="setUploadType('1')">
											                    <span>{{$.i18n.prop('create-single')}}</span>
											                  </li>
											                  <li @click="setUploadType('2')">
											                    <span>{{$.i18n.prop('create-multiple')}}</span>
											                  </li>
											                </ul>
											              </div>
											            </div>
										</div>
									</div>
									
									<template v-if="uploadType == 1">
										<div style="width:100%">
											<div class="col-12">
												<h4 class="sign__title i18n" name="create-uploadFiles">
													{{$.i18n.prop('create-uploadFiles')}}
												</h4>
											</div>
											<div class="col-12">
												<div class="sign__group">
													<div class="upload__img-container" @click="uploadCastingImg">
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
													<input v-model="title" :placeholder="$.i18n.prop('create-title')" type="text" class="sign__input i18n-input" select-name="create-pleaseEnterName" select-attr="placeholder">
												</div>
											</div>
											
											<div class="col-12">
												<div class="sign__group" style="position: relative;">
													<label class="sign__label i18n" name="create-desc">
														{{$.i18n.prop('create-desc')}}
													</label>
													<textarea v-model="desc" :disabled="descdisabled" name="text" :placeholder="$.i18n.prop('create-pleaseEnterDesc')" class="sign__textarea i18n-input" select-name="create-pleaseEnterDesc" select-attr="placeholder"></textarea>
													<span style="position: absolute;color: #FFFFFF;right: 20px;bottom: 0;">{{desclength}} / 500</span>
													
												</div>
											</div>
											<div class="col-12">
												<div class="sign__group sign__group-container">
													<label class="sign__label i18n" name="common-price">
														{{$.i18n.prop('common-price')}}
													</label>
													<div class="sign__input-container">
														<input v-model="price" :placeholder="$.i18n.prop('create-pleaseEnterPrice')" maxlength="8" type="number" class="sign__input i18n-input" select-name="create-pleaseEnterPrice" select-attr="placeholder">
														<span>ETH</span>
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
										</div>
									</template>
									<template v-if="uploadType == 2">
										<div style="width:100%">
											<div class="col-12">
												<h4 class="sign__title i18n" name="create-uploadFiles">
													{{$.i18n.prop('create-uploadFiles')}}
												</h4>
											</div>
											<div class="col-12">
												<div class="sign__group" style="flex-direction: column;align-items: flex-start">
													<div class="upload__img-container" @click="uploadMultipleImg">
														<div class="upload__img">
															<template v-if="imgArr.length == 0">
																<div v-if="!isloading" class="no_image" >
																	<img src="img/components/upload-icon.svg"/>
																	<span class="i18n" name="create-uploadImage">
																		{{$.i18n.prop('create-uploadImages')}}
																	</span>
																</div>
			
																<div v-else class="no_image">
																	<img class="loading" src="img/common/loading.svg"/>
																	<span>{{uploadIndex}} / {{fileLength}}</span>
																</div>
															</template>
															<template v-else>
																<div v-if="!isSuccess" class="no_image" >
																	<img src="img/components/upload-icon-failed.svg"/>
																	<span class="i18n" name="create-failed">
																		{{$.i18n.prop('create-failed')}}
																	</span>
																</div>
																<div v-if="isSuccess" class="no_image" >
																	<img src="img/components/upload-icon-completed.svg"/>
																	<span class="i18n" name="create-completed">
																		{{$.i18n.prop('create-completed')}}
																		
																	</span>
																</div>
																
															</template>
														</div>
														
													</div>
													
													<div style="color:#ffffff;margin-top:20px">{{$.i18n.prop('create-info')}}</div>
												</div>
											</div>
											<div class="col-12">
												<div class="sign__group sign__group-container">
													<label class="sign__label i18n" name="common-price">
														{{$.i18n.prop('common-price')}}
													</label>
													<div class="sign__input-container">
														<input v-model="price" maxlength="8" :placeholder="$.i18n.prop('create-pleaseEnterPrice')" type="number" class="sign__input i18n-input" select-name="create-pleaseEnterPrice" select-attr="placeholder">
														<span>ETH</span>
													</div>
												</div>
											</div>
										</div>
									</template>
					
									<div class="col-12" v-show="uploadType">
										<div class="row">
											<div class="col-3">
												<button class="sign__btn i18n" name="create-create" type="button" @click="mintHandler">
													{{$.i18n.prop('create-create')}}
												</button>
											</div>
											<div class="col-3" v-if="!collectionUriIn">
												<button class="sign__btn i18n" name="create-prev" type="button" @click="changeStep(2)">
													{{$.i18n.prop('create-prev')}}
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
			multipleDenied: true,
			uploadType: null,
			imgUrl: '',
			imgArr: [],
			cover: '',
			title: '',
			desc: '',
			price: '',
			royaltRate: 4,
			category: '',
			isloading: false,
			categorylist: [],
			signInfo: '',
			collectionUriIn: this.collectionuri,
			source: 1,
			coverInfo: {},
			isSuccess: false,
			uploadIndex:0,
			fileLength: 0
		}
	},
	props: ['collectionuri'],
	watch: {
		signInfo(newval, oldval) {
			if (newval != oldval && newval) {
				if (this.uploadType == '1') {
					if (!this.collectionUriIn) {
						this.saveCoverInfo()
					} else {
						this.castingSub()
					}
				} else {
					if (!this.collectionUriIn) {
						this.multipleCastingSub()
					} else {
						this.multipleCastingInfo()
					}

				}
			}
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
			},
			uploadTypeShow() {
			        let uploadTypeShow = '';
			        if (this.uploadType == 'single') {
			          uploadTypeShow = $.i18n.prop('create-single')
			        } else if (this.uploadType == 'multiple') {
						uploadTypeShow = $.i18n.prop('create-multiple')
			          // uploadTypeShow = this.$t('create.multiple')
			        }
			        return uploadTypeShow;
			      }
	},
	created() {
		this.getCategoryList();
		this.getPression();
		console.log(this.collectionUriIn)
	},
	methods: {
		setUploadType() {
			console.log(arguments[0])
		        this.uploadType = arguments[0];
		        this.isSelect = false;
		        this.uploadTypeChange()
		      },
		
		uploadTypeChange() {
		        this.cover = ''
		        this.coverloading = {
		          status: 0,// 0未选择图片之前 1开始上传 2上传失败 3上传成功
		          msg: ''
		        };
		
		        this.uploadIndex = 0;
		        this.fileLength = 0;
		        this.multipleCover = [];
		        this.itemname = '';//filename
		        this.description = '';//描述
		        this.price = ''//价格
		      },
		getCategoryList() {
			categoryList().then(res => {
				console.log(res)
				this.categorylist = res.list;
				if (res.list.length > 0) {
					this.category = res.list[0].id;
				}

				if (this.collectionUriIn) {
					getCollection({
						uri: this.collectionUriIn
					}).then(res => {
						const info = {
							category: res.category,
							royaltRate: res.royalt_rate,
							cover: res.cover,
							title: res.title,
							description: res.description
						}
						this.coverInfo = info;
					})
				} else {
					const coverInfo = JSON.parse(localStorage.getItem('coverInfo'));
					console.log(coverInfo)
					this.coverInfo = coverInfo
					this.category = coverInfo.category;
					this.cover = coverInfo.cover;
					this.royaltRate = coverInfo.royaltRate;
				}




			}).catch(res => {
				console.log(res)
			})
		},
		getPression() {
			isUpload().then(res => {
				if (res.bol) {
					this.multipleDenied = false
				} else {
					this.uploadType = 1
				}
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
					})
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
					message: $.i18n.prop('toast-deductionFailed'),
					duration: 1500
				});
				this.signInfo = '';
			}
		},
		//单个
		uploadCastingImg() {
			const uploadImg = new Upload({
				limit: 1, // 单次上传数量限制，默认9
				maxSize: 50 * 1024 * 1024, // 大小限制，公众号环境上传(env传wechat时)没有，默认10M
			});
			uploadImg.upload('', () => {
				this.isloading = true
			}).then(res => {
				this.isloading = false
				this.imgUrl = res.serverIds;
			}).catch((res) => {
				console.log(res)
			})
		},
		//duoge
		uploadMultipleImg() {
			const uploadImg = new Upload({
				limit: 100, // 单次上传数量限制，默认9
				maxSize: 50 * 1024 * 1024, // 大小限制，公众号环境上传(env传wechat时)没有，默认10M
				multiple: true,
				onProgress: (info) => {
					if (info.percent == 1) {
						this.uploadIndex++;
					}
				},
				startUpload: (len) => {
					this.uploadIndex = 0;
					this.fileLength = len;
					this.imgArr = []
				}
			});
			uploadImg.upload('', () => {
				this.isloading = true
			}).then(res => {
				this.isloading = false
				this.isSuccess = true;
				this.imgArr = res.serverIds;
				console.log('2323')
			}).catch((res) => {
				console.log(res)
				this.isSuccess = false;
			})
		},
		changeStep(step) {
			this.$emit('viewclick', step)
		},
		//提交封面信息
		async saveCoverInfo() {
			const coverInfo = this.coverInfo;
			await createCollection(coverInfo).then(res => {
				this.collectionUriIn = res.collectionUri;
				this.castingSub()
			})
		},
		//单个铸件
		castingSub() {
			const {
				collectionUriIn,
				cover,
				desc,
				price,
				sign,
				source,
				title
			} = this;
			createSubCasting({
				collectionUri: collectionUriIn,
				cover: cover,
				description: desc,
				price: +price,
				sign: sign,
				source: source,
				title: title
			}).then(() => {
				showToast({
					message: $.i18n.prop('create-createdSuccessfully'),
					duration: 5000,
					type: 'success',
					callback: () => {
						if (window.location.pathname.indexOf('profile') > -1) {
							this.$emit('viewclick')
						} else {
							window.location.href = "profile.html"
						}
					}
				});
			})
		},
		//新建合集并批量铸件
		multipleCastingSub() {
			const {
				coverInfo,
				imgArr,
				price
			} = this;
			worksUpload({
				images: imgArr,
				price: +price,
				...coverInfo
			}).then(() => {
				showToast({
					message: $.i18n.prop('create-createdSuccessfully'),
					duration: 5000,
					callback: () => {
						if (window.location.pathname.indexOf('profile') > -1) {
							this.$emit('viewclick')
						} else {
							window.location.href = "profile.html"
						}
					}
				});
			})
		},
		//在现有的合集里面批量铸件
		multipleCastingInfo() {
			const {
				collectionUriIn,
				imgArr,
				price
			} = this;
			worksUploadAdd({
				images: imgArr,
				price: +price,
				collectionUri: collectionUriIn
			}).then(() => {
				showToast({
					message: $.i18n.prop('create-createdSuccessfully'),
					duration: 5000,
					type: 'success',
					callback: () => {
						this.price = '';
						this.coverloading.status = 0;
						if (window.location.pathname.indexOf('profile') > -1) {
							this.$emit('viewclick')
						} else {
							window.location.href = "profile.html"
						}
					}
				});
			})
		}
	}
})

//  isUpload().then(res => {
//   isMultiple.value = res.bol;
//   if (!res.bol) {
//     uploadType.value = 'single';
//   }
// })
