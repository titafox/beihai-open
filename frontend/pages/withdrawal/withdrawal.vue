<template>
	<view class="custom_container">

		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="提现" :placeholder="true" class="nav_bar"
			:autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>

		</u-navbar>

		<u-modal title="请输入验证码" :show="show" closeOnClickOverlay @confirm="widthdrawalHandler">
			<u-code-input v-model="vericode" :maxlength="4"></u-code-input>
		</u-modal>

		<view class="card_container" v-if="cardList.length > 0">
			<u-text text="到账银行卡" color="#fff"></u-text>
			<u-cell-group :border="false">
				<u-cell @click="show1 = true" :title="curCard.bank" :label="curCard.no" :isLink="true"></u-cell>
			</u-cell-group>
		</view>

		<view class="card_container" v-else>
			<u-cell-group :border="false">
				<navigator url="/pages/cardEdit/cardEdit">
					<u-cell title="添加银行卡" :isLink="true"></u-cell>
				</navigator>
			</u-cell-group>
		</view>
		<u-gap height="20"></u-gap>

		<view v-if="cardList.length > 0">
			<u-text text="提现金额" color="#fff"></u-text>
			<u-gap height="40"></u-gap>
			<u-input v-model="amount" fontSize="28" border="bottom" type="number" :adjustPosition="false" clearable>

				<u-text text="¥" slot="prefix" size="20" margin="0 3px 0 0" type="tips"></u-text>
			</u-input>
			<u-gap height="10"></u-gap>
			<u-text :text="`可用余额 ¥ ${balance}`" v-show="!isOver" color="#fff"></u-text>
			<u-text text="输入金额超过可用余额" color="red" v-show="isOver"></u-text>

		</view>
		<u-gap height="30"></u-gap>
		<view class="btn_group" v-if="cardList.length > 0">
			<u-button color="linear-gradient(0deg, rgba(0, 0, 0, 0.2), rgba(0, 0, 0, 0.2)), linear-gradient(98.46deg, #4F75FF -0.18%, #6954FF 48.62%, #892CFF 100.99%, #892CFF 103.22%)" text="提现" size="large" :disabled="(amount > 0&& !isOver) ? false:true"
				@click="getVeriCode" shape="circle"></u-button>
				
		</view>

		<u-action-sheet :show="show1" @close="show1 = false" description="请选择提现银行卡" class="action_container">

			<u-list @scrolltolower="getCardList" height="200">
				<u-list-item v-for="card in cardList" :key="card.id">
					<u-cell :title="card.bank_office_name" :label="card.bank_num" @click="chooseBankHandler(card)">
						<u-icon color="#BC42FF" name="checkmark-circle-fill" slot="value"
							v-show="curCard.bank_id === card.id"></u-icon>
						<view slot="value" class="uncheck_radio" v-show="curCard.bank_id !== card.id"></view>

					</u-cell>
				</u-list-item>
			</u-list>

		</u-action-sheet>
	</view>
</template>

<script>
	import {
		accountBankListReq,
		withdrawalCode,
		withdrawal
	} from '../../api/index.js'
	export default {
		data() {
			return {
				page: 1,
				pageSize: 10,
				cardList: [],
				cur: 0,
				balance: '',
				amount: '',
				show: false,
				vericode: '',
				show1: false
			}
		},
		computed: {
			curCard() {
				let info = {
					bank: '',
					bank_id: '',
					no: ''
				}
				if (this.cardList.length > 0) {
					const {
						id,
						bank_num,
						bank_office_name
					} = this.cardList[this.cur];

					info = {
						bank: bank_office_name,
						bank_id: id,
						no: bank_num
					}
				}
				return info
			},
			isOver() {

				return this.amount > this.balance ? true : false;
			}
		},
		onShow() {
			this.page = 1;
			this.getCardList()

			let accountInfo = uni.getStorageSync('account');
			accountInfo = JSON.parse(accountInfo);
			const {
				amount
			} = accountInfo;
			this.balance = amount;
		},
		methods: {
			getCardList() {
				const {
					page,
					pageSize
				} = this;
				accountBankListReq({
					page,
					page_size: pageSize
				}).then(res => {
					const {
						result
					} = res
					if(!result)return
					if (page === 1) {
						this.cardList = result;
					} else {
						this.cardList.splice(this.cardList.length, 0, ...result)
					}
					this.page++;
				})
			},
			getVeriCode() {
				withdrawalCode().then(res => {
					this.show = true;
				})
			},
			widthdrawalHandler() {
				const {
					amount,
					vericode
				} = this;
				if (!vericode) return;
				const {
					bank_id
				} = this.curCard;
				withdrawal({
					amount: parseFloat(amount),
					bank_id,
					code: vericode
				}).then(res => {
					this.show = false;
					uni.$u.toast('操作成功！');

					uni.navigateBack()
				})
			},
			chooseBankHandler(card) {
				const {
					cardList
				} = this;


				cardList.map((item, i) => {
					if (item.id === card.id) {
						this.cur = i;
						this.show1 = false;
						return;
					}
				})


			}
		}
	}
</script>

<style scoped>
	>>>.u-line {
		border-bottom-width: 0 !important;
	}

	>>>.u-action-sheet {
		text-align: left;
	}
	>>> .u-cell__title-text,>>>.u-cell__label,>>>.u-text__value--tips,>>>.u-button__text{
		color:#fff;
	}

	>>>.u-list .u-cell__body {
		padding: 10px 25px;
	}

	.uncheck_radio {
		width: 12px;
		height: 12px;
		border: 1px solid #cfcfcf;
		border-radius: 50%;
	}
	
	.action_container >>>.u-transition .u-popup__content{
		background-color: #21005D;
	}
	
	>>>.u-action-sheet__description {
		color: #fff;
	}
</style>
