<template>
	<view>
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#ffffff00" class="nav_bar" :autoBack="true"
			leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		<view class="welcome_container">
			<u-image src="/static/welcome/welcome.png" width="100%" height="100%" :fade="true" duration="450"
				mode="aspectFill">
			</u-image>
		</view>
		<view class="info_container">
			<u-text class="title_1" text="北海" color="#FFFFFF" size="57"></u-text>
			<u-text class="title_1" text="NORTH" color="#FFFFFF" size="57"></u-text>
			<u-gap height="48"></u-gap>
			<u-text class="title_2" text="在北海·NORTH，由社区一起开启一扇全新的" color="#D0BCFF" size="14"></u-text>
			<u-text class="title_2" text="数字艺术大门。" color="#D0BCFF" size="14"></u-text>
		</view>
		<u-gap height="68"></u-gap>
		<view style="position: relative;width: 60px;margin: auto;"  @click="link2update">
			<u-text v-if="version" class="about_btn" :text="`V${version}`" align="center" size="11" color="#fff">
			</u-text>
			<u-badge v-show="isNew" :isDot="true" type="error" :absolute="true" :offset="[-2,-2]"></u-badge>
		</view>
		<u-gap height="10"></u-gap>
		<u-text v-show="isNew" text="有新版本，点击更新" align="center" size="11" color="rgba(255, 255, 255, 0.4)">
		</u-text> 
		
		<u-gap height="30"></u-gap>
	</view>
</template>

<script>
	import {
		versionReq
	} from '@/api/index.js'
	export default {

		data() {
			return {
				isNew: false,
				version: null
			}
		},
		onShow() {
			this.checkUpdate()
		},
		methods: {
			checkUpdate() {
				plus.runtime.getProperty(plus.runtime.appid, (inf) =>{
					this.getInfo(inf)

				});
			},
			getInfo(inf){
				const {
					version,
					versionCode
				} = inf;
				this.version = version
				versionReq().then(res => {
					const {
						newversion,
						enforce
					} = res;
					// enforce 0 不强制
					if (newversion > versionCode) {
						this.isNew = true
					}
				}).catch(err => {
					console.log(err)
				})
			},
			link2update(){
				if(this.isNew){
					const url = 'https://download.mareboreum.com'
					plus.runtime.openURL(url, () => {
						uni.$u.toast('打开下载页面失败！');
					});
				}
				
			}
		}
	}
</script>

<style>
	.welcome_container>>>.u-transition {
		width: 100vw;
		height: 700rpx;
	}

	.info_container {
		padding-top: 234rpx;
		padding-left: 106rpx;
	}


	.title_1>>>span {
		font-weight: 200;
		line-height: 64px;
		letter-spacing: -0.25px;
	}



	.title_2>>>span {
		font-weight: 500;
		line-height: 20px;
		letter-spacing: 0.1px;
	}

	.about_btn {
		background-color: #1B2374;
		flex: 0 0 auto;
		width: auto;
		padding: 10rpx 20rpx;
		border-radius: 8px;
	}
</style>
