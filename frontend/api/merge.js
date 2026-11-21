import request from '../utils/request.js'

export const getCompositeList = () => request({
	url: 'v1/nft/composite/list',
	method: 'GET'
})

export const getCompositeMaterial = (data) => request({
	url: 'v1/nft/composite/material',
	method: 'GET',
	data
})

export const doComposite = (data) => request({
	url: 'v1/nft/composite/start',
	data
})