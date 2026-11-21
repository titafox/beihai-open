//合集封面

Vue.component('collection-cover', {
	template: `
		<div class="collection-cover">
			<div class="tab-content">
				<div class="tab-pane fade show active" role="tabpanel">
					<div class="row row--grid">
						<div class="col-12">
							<form action="#" class="sign__form sign__form--profile">
								<div class="row">
									<div class="col-12">
										<h4 class="sign__title i18n" name="create-uploadCover">
											{{$.i18n.prop('create-uploadCover')}}
										</h4>
									</div>
									<div class="col-12">
										<div class="sign__group">
											<div class="upload__img-container" @click="uploadCoverImg">
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
											<input v-model="title" :placeholder="$.i18n.prop('create-pleaseEnterName')" type="text" class="sign__input i18n-input" select-name="create-pleaseEnterName" select-attr="placeholder">
										</div>
									</div>
									
									<div class="col-12">
										<div class="sign__group"  style="position:relative">
											<label class="sign__label i18n" name="create-desc">
												{{$.i18n.prop('create-desc')}}
											</label>
											<textarea v-model="desc" :disabled="descdisabled" :placeholder="$.i18n.prop('create-pleaseEnterDesc')" name="text" class="sign__textarea i18n-input" select-name="create-pleaseEnterDesc" select-attr="placeholder"></textarea>
											<span style="position: absolute;color: #FFFFFF;right: 20px;bottom: 0;">{{desclength}} / 500</span>
										
										</div>
									</div>
									
									<div class="col-12">
										<div class="sign__group">
											<label class="sign__label i18n" name="create-classification" >
												{{$.i18n.prop('create-classification')}}
											</label>
											<select v-if="!collectionuri" name="select" class="sign__select" v-model="category">
												<option v-for="(item,i) in categorylist" :key="item.id" :value="item.id" >{{item.categoryName}}</option>
											</select>
											
											<div v-if="collectionuri" class="sign__input-container">
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
											<select v-if="!collectionuri" name="select" class="sign__select" v-model="royaltRate">
												<option value="4">4%</option>
											    <option value="5">5%</option>
											    <option value="6">6%</option>
											    <option value="7">7%</option>
											    <option value="8">8%</option>
											    <option value="9">9%</option>
											    <option value="10">10%</option>
											</select>
											
											<div v-if="collectionuri" class="sign__input-container">
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
										<button  v-if="!collectionuri" class="sign__btn i18n" name="create-next" type="button" @click="nextHandler">
											{{$.i18n.prop('create-next')}}
										</button>
										
										<button  v-if="collectionuri" class="sign__btn i18n" name="create-save" type="button" @click="editCollectionInfo">
											{{$.i18n.prop('create-save')}}
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
			royaltRate: 4,
			category: '',
			isloading: false,
			categorylist: [],
			status: ''
		}
	},
	props: ['collectionuri'],
	created() {
		this.getCategoryList()
		
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
	methods: {
		getCategoryList() {
			categoryList().then(res => {
				this.categorylist = res.list
				if (res.list.length > 0) {
					this.category = res.list[0].id;
				}
				console.log(this.collectionuri)
				if(this.collectionuri){
					this.getCollectionInfo()
				}
			}).catch(res => {
				console.log(res.list)
			})
		},
		uploadCoverImg() {

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
		nextHandler() {
			const {
				category,
				imgUrl,
				desc,
				royaltRate,
				title
			} = this
			if (!imgUrl || !royaltRate || !title) {
				showModal({
					type: 'error',
					message: $.i18n.prop('create-pleaseEnterCompleteInformation'),
				});
				return
			}
			console.log(royaltRate)
			const coverInfo = {
				category: category,
				cover: imgUrl,
				description: desc,
				royaltRate: royaltRate,
				title: title
			}

			localStorage.setItem('coverInfo', JSON.stringify(coverInfo));
			this.$emit('viewclick', 4)
		},
		editCollectionInfo() {
			const {
				imgUrl,
				desc,
				royaltRate,
				title,
				status,
				collectionuri
			} = this
			if (!imgUrl || !title) {
				showModal({
					type: 'error',
					message: $.i18n.prop('create-pleaseEnterCompleteInformation'),
					
				});
				return
			}
			editCollection({
				cover: imgUrl,
				description: desc,
				royaltRate: +royaltRate,
				title: title,
				status: status,
				uri: collectionuri
			}).then(() => {
				showToast({
					message: $.i18n.prop('create-createdSuccessfully'),
					type: 'success',
					duration: 5000
				});
				this.$emit('viewclick')
			})
		},
		getCollectionInfo() {
			getCollection({
				uri: this.collectionuri
			}).then(res => {
				this.imgUrl = res.cover;
				this.title = res.title;
				this.desc = res.description;
				this.royaltRate = res.royalt_rate;
				this.category = res.category;
				this.status= res.status;
			})
		}
	}
})
