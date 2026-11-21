<template>
	<view class="custom_container">
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="基础设置" :placeholder="true" class="nav_bar"
			:autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		<view class="profile_header">
			<u-avatar size="120" :src="userInfo.avatar" class="avatar"></u-avatar>
			<u-upload multiple class="edit_icon" :previewImage="false" :fileList="fileList1" @afterRead="afterRead"
				name="1" :maxCount="1" width="32" height="32">
				<u-icon name="/static/profile_icon/upload.png"></u-icon>
			</u-upload>
		</view>



		<u-gap height="15"></u-gap>

		<u-form labelPosition="top" :model="userInfo" ref="form">
			<u-form-item label="用户名" labelWidth="auto">
				<u-input :adjustPosition="false" v-model="userInfo.nickname" placeholder="输入你的用户名"></u-input>
			</u-form-item>
			
			<u-form-item label="个人简介" prop="desc" labelWidth="auto">
				<u-textarea  v-model="userInfo.desc" placeholder="最多可以输入500个字符" confirmType="done">
				</u-textarea>
			</u-form-item>
		</u-form>
		<u-gap height="45"></u-gap>

		<view class="btn_group">
			<u-button shape="circle" color="linear-gradient(0deg, rgba(0, 0, 0, 0.2), rgba(0, 0, 0, 0.2)), linear-gradient(98.46deg, #4F75FF -0.18%, #6954FF 48.62%, #892CFF 100.99%, #892CFF 103.22%)" text="保存信息" size="large" @click="updatePersonalInfo"></u-button>
		</view>
	</view>
</template>

<script>
	import {
		avatarUploadReq,
		accountUpdateReq
	} from '../../api/index.js';
	export default {
		data() {
			return {
				userInfo: {
					nickname: '',
					avatar: '',
					desc:''
				},
				fileList1: [],
			};
		},
		onLoad() {
			let accountInfo = uni.getStorageSync('account')
			accountInfo = JSON.parse(accountInfo)
			const {
				nickname,
				avatar,
				description
			} = accountInfo;
			this.userInfo = {
				nickname,
				avatar,
				desc:description
			}
		},

		methods: {
			deletePic(event) {
				this.fileList1.splice(event.index, 1);
			},
			async afterRead(event) {
				// const result = await this.uploadFilePromise(event.file.url);

				let lists = [].concat(event.file)
				let fileListLen = this[`fileList${event.name}`].length

				// lists.map((item) => {
				// 	this[`fileList${event.name}`].push({
				// 		...item,
				// 		status: 'uploading',
				// 		message: '上传中'
				// 	})
				// })
				for (let i = 0; i < lists.length; i++) {
					const result = await this.uploadFilePromise(lists[i].url)
					// let item = this[`fileList${event.name}`][fileListLen]
					// this[`fileList${event.name}`].splice(fileListLen, 1, Object.assign(item, {
					// 	status: 'success',
					// 	message: '',
					// 	url: result
					// }))
					fileListLen++
				}

			},
			uploadFilePromise(url) {
				avatarUploadReq(url).then(res => {

					this.userInfo.avatar = res;
				})
			},
			updatePersonalInfo() {
				const {
					nickname,
					avatar,
					desc
				} = this.userInfo;
				accountUpdateReq({
					nickname,
					avatar,
					desc
				}).then(() => {
					uni.$u.toast('保存成功');
					setTimeout(()=>{
						uni.navigateBack()
					},2000)
				})
			}

		}
	};
</script>

<style scoped>
	.profile_header {
		position: relative;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.avatar_container {
		width: 120px;
		height: 120px;
		position: absolute;
		left: 0;
		top: 120px;
		right: 0;
		bottom: 0;
		margin: auto;
		background-color: #fefefe;
		border-radius: 50%;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.profile_container .author_name {
		font-weight: 500;
		line-height: 26px;
		letter-spacing: 0.005em;
	}

	.edit_icon {
		background-color: rgba(20, 15, 15, 0.9);
		width: 120px;
		height: 120px;
		border-radius: 50%;
		display: flex;
		align-items: center;
		justify-content: center;
		position: absolute;
		top: 0;
		right: 0;
		bottom: 0;
		left: 0;
		margin: auto;

	}

	.edit_icon>>>.u-upload__wrap {
		flex-shrink: 0;
		flex-grow: 0;
		flex-basis: auto;
	}
	
	
</style>
