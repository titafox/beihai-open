<template>
	<view class="custom_container">
		<u-navbar :titleStyle="{'color': '#fff'}" bgColor="#21005D" title="Artist" :placeholder="true" class="nav_bar"
			:autoBack="true" leftIconColor="#151517">
			<u-icon name="arrow-leftward" color="#fff" slot="left"></u-icon>
		</u-navbar>

		<u-text text="北海的优秀" size="36" color="#ffffff"></u-text>
		<u-text text="艺术家们" size="36" color="#ffffff"></u-text>
		<u-gap height="15"></u-gap>

		<u-text text="艺术的脚步从未停止，为了爱与远方，御而不舍。" color="#D0BCFF" size="12"></u-text>

		<u-gap height="40"></u-gap>

		<u-list @scrolltolower="scrolltolower" class="artist_list">
			<u-list-item v-for="artist in list" :key="artist.id">
				<view class="author_card_container">
					<view class="avatar_container">
						<u-avatar :src="artist.avatar" size="65"></u-avatar>
					</view>
					<u-gap height="50"></u-gap>
					<u-text :lines="2" :text="artist.nickname" color="#EADDFF" size="12"></u-text>
					<u-gap height="11"></u-gap>
					<u-text :text="`${artist.works_count} Items`" color="#7F67BE" size="11"></u-text>
				</view>
			</u-list-item>
			<u-divider text="没有更多了..." :hairline="true" v-show="pageSize*page >= count"></u-divider>
		</u-list>

	</view>
</template>

<script>
	import {
		artistsReq
	} from '../../api/index.js'
	export default {
		data() {
			return {
				page: 1,
				pageSize: 20,
				isReq: false,
				count: 21,
				list: []
			}
		},
		onLoad() {
			this.getArtistList()
		},
		methods: {
			getArtistList() {
				artistsReq().then(res => {
					this.list = res.result;

				})
			},
			scrolltolower() {

			}
		}
	}
</script>

<style scoped>
	.author_card_container {
		background: #381E72;
		border-radius: 12px;
		padding: 24rpx 36rpx;
		display: flex;
		flex-direction: column;
		align-items: center;
		margin-bottom: 36rpx;
		position: relative;
		height: 228rpx;
	}

	.avatar_container {
		position: absolute;
		top: -10px;
	}

	>>>.u-list-item {
		width: 30%;
	}

	.artist_list>>>.uni-scroll-view-content>view {
		display: flex;
		flex-wrap: wrap;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		padding-top: 40rpx;
	}

	/* .author_card_container>>>.u-list-item {} */
</style>
