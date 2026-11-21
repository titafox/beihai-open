<template>
	<view class="layout-container">
		<u-navbar
			:titleStyle="{'color': '#fff'}" 
			bgColor="#21005D" 
			:title="title"
			:placeholder="true" 
			class="nav_bar"
			:autoBack="true" 
			leftIconColor="#151517"
		>
			<u-icon name="arrow-leftward" color="#fff" slot="left" />
		</u-navbar>
		<u-gap height="20" />
		<solt></solt>
	</view>
</template>

<script>
	export default {
		name: 'PageLayout',
		props: {
			title: {
				type: String,
				default: '标题'
			}
		}
	}
</script>

<style lang="less" scoped>
	.layout-container {
		padding: 48rpx 28rpx calc(env(safe-area-inset-bottom) / 2 + 24rpx);
	}
</style>