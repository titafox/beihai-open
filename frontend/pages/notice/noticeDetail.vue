<template>
	<view class="custom_container">
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="公告" :placeholder="true" class="nav_bar"
			:autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>
		<u-text :text="notice.title" color="#fff" size="24"></u-text>
		<view v-html="notice.content">

		</view>
	</view>
</template>

<script>
	import {
		noticeDetailReq
	} from '../../api/index.js'
	export default {
		data() {
			return {
				id: null,
				notice: {
					content: '',
					title: ''
				}
			}
		},
		onLoad(options) {
			this.id = options.id;
			this.getNoticeDetail()
		},
		methods: {
			getNoticeDetail() {
				const {
					id
				} = this;
				noticeDetailReq({
					id
				}).then(res => {
					const {
						content,
						title
					} = res
					this.notice = {
						title,
						content
					}

				})
			}
		}
	}
</script>

<style>

</style>
