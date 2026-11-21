<template>
	<view>
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="手机号" :placeholder="true" class="nav_bar" :autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>

		<view class="custom_container">
			<u-form labelPosition="top" :model="phoneInfo" :rules="rules" ref="form">

				<u-form-item label="原手机号" labelWidth="auto">
					<!-- <u-input :adjustPosition="false" v-model="phoneInfo.originPhone" placeholder="" :disabled="true">
					</u-input> -->
					
					<u-text :text="phoneInfo.originPhone" size="14" color="#D0BCFF"></u-text>
				</u-form-item>
				<u-form-item label="手机号码" prop="phone" labelWidth="auto">
					<u-input :adjustPosition="false" v-model="phoneInfo.phone" placeholder="输入你的手机号码">
						<template slot="suffix">
							<u-code start-text="发送验证码" ref="uCode" @change="codeChange" seconds="60"
								changeText="X秒重新获取"></u-code>
							<u-text @tap="getCode" color="#FF3838" size="13" :text="tips" class="u-page__code-text">
								{{tips}}
							</u-text>
						</template>
					</u-input>
				</u-form-item>
				<u-form-item label="验证码" prop="veriCode" labelWidth="auto">
					<u-code-input mode="line" v-model="phoneInfo.veriCode" maxlength="6" dot size="58" :adjustPosition="false">
					</u-code-input>
				</u-form-item>
			</u-form>
			<u-gap height="27"></u-gap>
			<view class="btn_group">
				<u-button shape="circle" color="linear-gradient(0deg, rgba(0, 0, 0, 0.2), rgba(0, 0, 0, 0.2)), linear-gradient(98.46deg, #4F75FF -0.18%, #6954FF 48.62%, #892CFF 100.99%, #892CFF 103.22%)" text="更新手机号码" size="large" @click="updatePhoneHandler"></u-button>
			</view>
		</view>
	</view>
</template>

<script>
	import {
		updateVeriCodeReq,
		updatePhoneReq
	} from '../../api/index.js'
	export default {
		data() {
			return {
				phoneInfo: {
					originPhone: '',
					phone: '',
					veriCode: '',
				},
				rules: {
					phone: [{
							type: 'string',
							required: true,
							message: '请填写正确的手机号',
							trigger: ['blur', 'change']
						},
						{
							pattern: (/^1[3456789]\d{9}$/),
							transform(value) {
								return String(value);
							},
							message: '请填写正确的手机号'
						},
					],
					veriCode: {
						type: 'string',
						required: true,
						len: 6,
						message: '请填写6位验证码',
						trigger: ['blur']
					},
				},
				tips: '',
			};
		},
		onShow() {
			let accountInfo = uni.getStorageSync('account');
			accountInfo = JSON.parse(accountInfo);
			const {
				phone,
			} = accountInfo;
			this.phoneInfo = {
				...this.phoneInfo,
				originPhone: phone
			}
		},
		methods: {
			codeChange(text) {
				this.tips = text;
			},
			getCode() {
				const {
					phone
				} = this.phoneInfo;
				const reg = new RegExp(/^1[3456789]\d{9}$/);

				if (!reg.test(phone)) {
					uni.$u.toast('请输入正确的手机号');
					return false;
				}

				if (this.$refs.uCode.canGetCode) {
					uni.showLoading({
						title: '正在获取验证码'
					})

					updateVeriCodeReq({
						phone
					}).then(res => {
						uni.hideLoading();
						uni.$u.toast('验证码已发送');
						this.$refs.uCode.start();
					}).catch(err => {

					})
				} else {
					uni.$u.toast('倒计时结束后再发送');
				}
			},
			updatePhoneHandler() {
				this.$refs.form.validate().then(res => {
					const {
						phone,
						veriCode
					} = this.phoneInfo;
					updatePhoneReq({
						phone,
						code: veriCode
					}).then(() => {
						uni.navigateBack()
					})
				}).catch(errors => {})
			}
		}
	};
</script>

<style scoped>
	>>>.u-form-item__body__right__content__slot {
		justify-content: space-around;
	}
</style>
