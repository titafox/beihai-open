/* 1.*/

var vm = new Vue({
	el: '#content',
	data() {
		return {
			baseInfo: {
				paintingWidth: 24, //成品画作的宽高为 24*24
				paintingCanvasWidth: 816, // 创作画布的宽高为816*816
				previewCanvasWidth: 480, //预览图片的宽高为480*480
				scaleCanvasWidth: 1200

			},
			color: '#59c7f9',
			alpha: 1,
			suckerCanvas: null,
			suckerArea: [],
			isSucking: false,
			currentTool: 'pencil',
			layersArray: [
				// {
				//   name: '',
				//   id: '',
				//   flod: false,// 折叠 false 即展开
				//   isChosen: false, //是否选中
				//   images: {
				//     name: '',
				//     id: '',
				//     visible: true, // 可见
				//     isChosen: false //是否选中
				//   }
				// }
			],
			isChosen: false,
			editInfo: {
				name: '',
				placeholder: '',
				isEdit: false
			},
			copyInfo: {
				isEdit: false,
				origin: 1,
				change: 1,
			},
			editImageId: '',
			editLayerId: '',
			isDrawing: false,
			drawInfo: {
				pos: [],
				alpha: '',
				color: ''
			},
			zoomIndexInfo: [
				// 'imageId'
			], //层级关系 每次layersArray变化的时候重新计算

			frameInfo: {
				layersArray: [],
				canvasInfo: [],
				imageArray: []
			}, //每次操作的信息
			currentIndex: -1, //frameInfo.length -1
			paintingInfo: {
				// 'imageId': [
				// {
				// //     pos: [],
				// //     alpha: '',
				// //     color: []
				// //   }
				// ]
			}, //图片上的信息

			mousePos: [], //实时的鼠标位置
			newLayerId: 0,
			newImageId: 0,
			bgColor: [],
			imageZipKey: '',
			multiple: 8,
			loading: false,
			toastInfo: '',
			infoToggle: false,
			previewImage: '',
			previewToast: '',
			websocket: null,
			isShow: false,
		}
	},
	computed: {
		nameObj() {
			let nameObject = {
				layers: [],
				images: []
			};

			const layersArray = this.layersArray;
			let layers = [];
			let images = [];
			for (let i = 0; i < layersArray.length; i++) {
				layers.push(layersArray[i].name)
				for (let j = 0; j < layersArray[i].images.length; j++) {
					images.push(layersArray[i].images[j].name);
				}
			}
			nameObject = {
				layers: layers,
				images: images
			}
			return nameObject;
		},

	},
	watch: {
		mousePos(newval, oldval) {

			if (newval[0] == oldval[0] && newval[1] == oldval[1]) {
				return
			}

			const x = newval[0];
			const y = newval[1];
			const currentTool = this.currentTool;
			const editImageId = this.editImageId;
			const color = this.color;
			
			const alpha = this.alpha;
			const baseInfo = this.baseInfo
			const width = baseInfo.paintingCanvasWidth / baseInfo.paintingWidth;
			const zoomIndexInfo = this.zoomIndexInfo;


			const imageIndex = zoomIndexInfo.indexOf(editImageId);

			const index = x / width + y / width * baseInfo.paintingWidth; //绘图的位置
			
			
			
			if (currentTool == 'pencil') {
				//fill
				// const zoomIndexInfo = this.zoomIndexInfo;

				let paintingCanvas = document.getElementById("painting-canvas");
				let paintingCtx = paintingCanvas.getContext("2d");
				paintingCtx.fillStyle = color;
				paintingCtx.globalAlpha = alpha;



				paintingCtx.clearRect(x, y, width, width);
				paintingCtx.fillRect(x, y, width, width);




				let canvasBg = document.getElementById("canvasBg");
				let ctx = canvasBg.getContext('2d');
				ctx.fillStyle = '#ffffff';
				ctx.globalAlpha = 1;
				ctx.clearRect(x, y, width, width);
				ctx.fillRect(x, y, width, width);



				let drawInfo = this.drawInfo;
				let pos = drawInfo.pos;
				//获取位置
				if (pos.length === 0) {
					pos = [];
					for (let n = 0; n < baseInfo.paintingWidth * baseInfo.paintingWidth; n++) pos[n] = 0
				}

				pos.splice(index, 1, 1);
				drawInfo.pos = pos;
				drawInfo.color = this.color;
				drawInfo.alpha = this.alpha;
				drawInfo.type = 'fill';
				this.drawInfo = drawInfo;
			} else {
				//clear
				const index = x / width + y / width * baseInfo.paintingWidth;
				let paintingCanvas = document.getElementById("painting-canvas");
				let paintingCtx = paintingCanvas.getContext("2d");

				paintingCtx.clearRect(x, y, width, width);


				let canvasBg = document.getElementById("canvasBg");
				let ctx = canvasBg.getContext('2d');
				const bgColor = this.bgColor;

				ctx.fillStyle = bgColor[index];
				ctx.globalAlpha = 1;
				ctx.clearRect(x, y, width, width);
				ctx.fillRect(x, y, width, width);



				let drawInfo = this.drawInfo;
				let pos = drawInfo.pos;

				//获取位置
				if (pos.length === 0) {
					pos = [];
					for (let n = 0; n < baseInfo.paintingWidth * baseInfo.paintingWidth; n++) pos[n] = 0
				}

				pos.splice(index, 1, 1);

				drawInfo.pos = pos;
				drawInfo.color = '';
				drawInfo.alpha = '';
				drawInfo.type = 'clear';
				this.drawInfo = drawInfo;
			}

		},
		isDrawing(newval) {
			//绘制结束
			if (!newval) {
				const editImageId = this.editImageId;
				const drawInfo = this.drawInfo;
				const pos = JSON.parse(
					JSON.stringify(drawInfo.pos)
				);
				const option = {
					id: editImageId,
					pos: pos,
					color: drawInfo.color,
					alpha: drawInfo.alpha,
					type: drawInfo.type
				}

				let paintingInfo = JSON.parse(JSON.stringify(this.paintingInfo));

				if (!paintingInfo[editImageId]) {
					paintingInfo[editImageId] = []
				}

				for (let i = 0; i < paintingInfo[editImageId].length; i++) {
					let pos = paintingInfo[editImageId][i].pos;
					for (let j = 0; j < pos.length; j++) {
						if (drawInfo.pos[j]) pos[j] = 0
					}
					paintingInfo[editImageId][i].pos = pos;
				}

				if (drawInfo.type == 'fill') {
					paintingInfo[editImageId].push({
						pos: drawInfo.pos,
						color: drawInfo.color,
						alpha: drawInfo.alpha
					})
				}
				this.paintingInfo = paintingInfo;
				console.log(paintingInfo)
				this.drawInfo = {
					pos: [],
					alpha: ' ',
					color: '',
					type: ''
				}

				this.setPerviewImage();


				this.setFrameInfo()

			}


		},
		layersArray(newval) {
			let zoomIndexInfo = [];
			// 计算图片层级
			const layersArray = this.layersArray;



			for (let i = 0; i < layersArray.length; i++) {
				for (let j = 0; j < layersArray[i].images.length; j++) {

					if (layersArray[i].images[j].visible) {
						zoomIndexInfo.push(layersArray[i].images[j].id)
					}

				}
			}
			this.zoomIndexInfo = zoomIndexInfo;


			this.setPerviewImage();



		},
		editImageId(newval, oldval) {
			this.setEmptyCanvas()
			if (newval) {
				console.log('reset painting')
				const paintingInfo = this.paintingInfo;
				console.log(paintingInfo)
				const paintingArray = paintingInfo[newval];
				if (!paintingArray) return
				let colorArray = [];
				let posArray = [];
				let alphaArray = [];
				const baseInfo = this.baseInfo;
				let paintingCanvas = document.getElementById('painting-canvas');
				let paintingCtx = paintingCanvas.getContext('2d');
				let isWhite = false;

				const width = baseInfo.paintingCanvasWidth / baseInfo.paintingWidth; // 816/24
				for (let i = 0; i < baseInfo.paintingWidth; i++) {
					isWhite = !isWhite;
					for (let j = 0; j < baseInfo.paintingWidth; j++) {
						isWhite = !isWhite;
						paintingCtx.fillStyle = isWhite ? "#A1A8A1" : "#FFFFFF";
						paintingCtx.fillRect(width * i, width * j, width, width);
					}
				}

				for (let n = 0; n < baseInfo.paintingWidth * baseInfo.paintingWidth; n++) {
					posArray[n] = 0
					colorArray[n] = '';
					alphaArray[n] = 1;
				}


				if (paintingArray) {
					for (let i = 0; i < paintingArray.length; i++) {
						let alpha = paintingArray[i].alpha;
						let color = paintingArray[i].color;
						let pos = paintingArray[i].pos;
						for (let j = 0; j < pos.length; j++) {
							if (pos[j]) {
								posArray.splice(j, 1, pos[j])
								colorArray.splice(j, 1, color);
								alphaArray.splice(j, 1, alpha);

							}
						}
					}
				}

				let canvasBg = document.getElementById("canvasBg");
				let ctx = canvasBg.getContext("2d");

				for (let i = 0; i < baseInfo.paintingWidth; i++) {
					for (let j = 0; j < baseInfo.paintingWidth; j++) {
						if (posArray[i + j * baseInfo.paintingWidth]) {
							paintingCtx.fillStyle = colorArray[i + j * baseInfo.paintingWidth];
							paintingCtx.globalAlpha = alphaArray[i + j * baseInfo.paintingWidth];
							paintingCtx.clearRect(width * i, width * j, width, width)
							paintingCtx.fillRect(width * i, width * j, width, width)

							ctx.fillStyle = '#fff';
							ctx.globalAlpha = 1;
							ctx.clearRect(width * i, width * j, width, width)
							ctx.fillRect(width * i, width * j, width, width)
						}

					}
				}

			}
		},
		
	},
	mounted() {
		console.log(window.screen.availHeight - 70)
		if(window.screen.availHeight - 70 < 1000){
			document.getElementById('canvasBg').setAttribute('height',600);
			document.getElementById('canvasBg').setAttribute('width',600);
			document.getElementById('painting-canvas').setAttribute('width',600)
			document.getElementById('painting-canvas').setAttribute('height',600)
			document.getElementById('scale-canvas').setAttribute('width',720)
			document.getElementById('scale-canvas').setAttribute('height',720)
			
			
			document.getElementById('preview-canvas').setAttribute('height',384)
			document.getElementById('preview-canvas').setAttribute('width',384)
			
			
			document.getElementById('circle-canvas').setAttribute('height',384)
			document.getElementById('circle-canvas').setAttribute('width',384)
		
			$('.preview-container .painting-panel').css({
				'transform':'scale(0.8)',
				"position": 'absolute',
				'top': '-40px',
				'left': '-35px'
			})
			
			
			$('.preview-title').css({
				'width':'384px'
			})
			
			// $('.preview-container').find('.col-12').eq(1).css({
			// 	"display": "flex",
			// 	"justify-content": "end"
			// })
			
			$('.preview-container').find('.col-12').eq(1).css({
				// "display": "flex"
			})
			
			$('#circle-canvas').css('left','15px');
			
			this.scale = 0.8
			
			this.baseInfo =  {
				paintingWidth: 24, //成品画作的宽高为 24*24
				paintingCanvasWidth: 600, // 创作画布的宽高为816*816
				previewCanvasWidth: 384, //预览图片的宽高为480*480
				scaleCanvasWidth: 720
			}
		}
		this.initCanvas();
		this.changeTool('pencil')
		this.addLayer('init')
	},
	unmounted() {
		//移除监听事件
	},
	methods: {
		setFrameInfo(){
			const layersArray = this.layersArray
			
			this.frameInfo.layersArray.push(JSON.parse(JSON.stringify(layersArray)));
			
			const editImageId = this.editImageId
			this.frameInfo.imageArray.push(editImageId)
			
			
			const paintingInfo  = this.paintingInfo
			this.frameInfo.canvasInfo.push(JSON.parse(JSON.stringify(paintingInfo)));
			
			//1.绘制完成 2.新增图片  3.删除图片 4.复制图片 5.重命名图片
			//1.新增图层 2.删除图层 3.复制图层 4.重命名图层
			
			if(this.frameInfo.imageArray.length > 20){
				const frameInfo = this.frameInfo;
				frameInfo.layersArray.splice(0,1)
				frameInfo.imageArray.splice(0,1)
				frameInfo.canvasInfo.splice(0,1)
				this.frameInfo = frameInfo
			} else {
				this.currentIndex = this.frameInfo.canvasInfo.length-1;
			}
			
			console.log(this.frameInfo.canvasInfo);
			console.log(this.currentIndex)
			
			
		},
		changeTool(tool) {
			this.currentTool = tool;
			let imgUrl = 'img/pages/eraser.svg'
			if (tool == 'pencil') {
				imgUrl = 'img/pages/pencil.svg'
			}

			document.getElementById('painting-canvas').style.cursor = `url(${imgUrl}) 0 32, default`

		},

		changeColor(color) {
			console.log(color)
			const {
				a
			} = color.rgba;
			this.color = color.hex;
			this.alpha = a;
			this.openSucker(false)
		},

		openSucker(isOpen) {
			console.log(isOpen)
			if (isOpen) {
				this.suckerCanvas = document.getElementById('preview-canvas')
				const width = this.baseInfo.previewCanvasWidth
				this.suckerArea = [0, 0, width, width];
				document.getElementById('circle-canvas').style.zIndex = -1
			} else {
				this.suckerCanvas = null
				document.getElementById('preview-canvas').style.cursor = 'default';
				this.changeTool('pencil')
				document.getElementById('circle-canvas').style.zIndex = 3;
			}
		},
		isVisible(imageId) {
			let layersArray = this.layersArray;
			let isVisible = false;
			for (let i = 0; i < layersArray.length; i++) {
				for (let j = 0; j < layersArray[i].images.length; j++) {
					if (imageId == layersArray[i].images[j].id) {
						isVisible = layersArray[i].images[j].visible;
					}
				}
			}
			return isVisible;
		},
		setEmptyCanvas() {
			let paintingCanvas = document.getElementById("painting-canvas");
			let paintingCtx = paintingCanvas.getContext("2d");
			let isWhite = false;
			const baseInfo = this.baseInfo;

			const width = baseInfo.paintingCanvasWidth / baseInfo.paintingWidth; // 816/24
			let canvasBg = document.getElementById("canvasBg");
			let ctx = canvasBg.getContext("2d");
			for (let i = 0; i < baseInfo.paintingWidth; i++) {
				isWhite = !isWhite;
				for (let j = 0; j < baseInfo.paintingWidth; j++) {
					isWhite = !isWhite;
					paintingCtx.fillStyle = isWhite ? "#A1A8A1" : "#FFFFFF";
					paintingCtx.clearRect(width * i, width * j, width, width);
					paintingCtx.fillRect(width * i, width * j, width, width);

					ctx.fillStyle = isWhite ? "#A1A8A1" : "#FFFFFF";
					ctx.clearRect(width * i, width * j, width, width);
					ctx.fillRect(width * i, width * j, width, width);

				}
			}
		},
		initCanvas() {
			let paintingCanvas = document.getElementById("painting-canvas");
			let paintingCtx = paintingCanvas.getContext("2d");
			let isWhite = false;
			const baseInfo = this.baseInfo;

			const width = baseInfo.paintingCanvasWidth / baseInfo.paintingWidth; // 816/24

			let canvas = document.getElementById("canvasBg");
			let ctx = canvas.getContext("2d");
			let bgColor = [];

			let scaleCanvas = document.getElementById('scale-canvas');
			let scaleCtx = scaleCanvas.getContext("2d");

			const scaleWidth = baseInfo.scaleCanvasWidth / baseInfo.paintingWidth;

			for (let i = 0; i < baseInfo.paintingWidth; i++) {
				isWhite = !isWhite;
				for (let j = 0; j < baseInfo.paintingWidth; j++) {
					isWhite = !isWhite;
					let color = isWhite ? "#A1A8A1" : "#FFFFFF";
					ctx.fillStyle = color;
					ctx.globalAlpha = 0;
					bgColor.push(color);
					ctx.fillRect(width * i, width * j, width, width);


					paintingCtx.fillStyle = isWhite ? "#A1A8A1" : "#FFFFFF";
					paintingCtx.fillRect(width * i, width * j, width, width);



					scaleCtx.fillStyle = isWhite ? "#A1A8A1" : "#FFFFFF";
					scaleCtx.fillRect(scaleWidth * i, scaleWidth * j, scaleWidth, scaleWidth);


				}
			}
			this.bgColor = bgColor;


			let previewCanvas = document.getElementById("preview-canvas");
			let previewCtx = previewCanvas.getContext("2d");
			previewCtx.globalAlpha = 0;
			previewCtx.fillStyle = "#FFFFFF";
			previewCtx.fillRect(0, 0, baseInfo.previewCanvasWidth, baseInfo.previewCanvasWidth);

			// paintingInfo bgGray bgWhite

			paintingCanvas.addEventListener('mousedown', (e) => {


				if (this.nameObj.layers.length == 0) {

					
					
					if(this.suckerCanvas){
						showModal({
							message:'Please intake color in the Preview Area!',
							type: 'error',
						})
					}else {
						showModal({
							message:'Please Add A New Layer!',
							type: 'error',
						})
					}
					return;
				}

				if (this.nameObj.images.length == 0) {
					
					if(this.suckerCanvas){
						showModal({
							message:'Please intake color in the Preview Area!',
							type: 'error',
						})
					}else {
						showModal({
							message:'Please Add A New Image!',
							type: 'error',
						})
					}
					return;
				}

				if (this.editImageId == '') {

					if(this.suckerCanvas){
						showModal({
							message:'Please intake color in the Preview Area!',
							type: 'error',
						})
					}else {
						showModal({
							message:'Please Choose A New Image!',
							type: 'error',
						})
					}
					return;
				}

				if (this.suckerCanvas) {
					// 取色器取色中...
					showModal({
						message:'Please intake color in the Preview Area!',
						type: 'error',
					})
					return;
				}



				this.isDrawing = true;
				//开始绘图
			})

	
			paintingCanvas.addEventListener('mousemove', (e) => {
				const cRect = paintingCanvas.getBoundingClientRect()
				var canvasX = Math.round(e.clientX - cRect.left);
				var canvasY = Math.round(e.clientY - cRect.top);
				const baseInfo = this.baseInfo;
				const paintingCanvasWidth = baseInfo.paintingCanvasWidth;
				
				const x = parseInt(canvasX / width) * width
				const y = parseInt(canvasY / width) * width
				
				
				const color = this.color;
				const alpha = this.alpha;
				let circleCanvas = document.getElementById("circle-canvas");
				let circleCtx = circleCanvas.getContext("2d");
				
				
				const prewidth = baseInfo.previewCanvasWidth/ baseInfo.paintingWidth
				
				const prewX  = baseInfo.previewCanvasWidth* x/baseInfo.paintingCanvasWidth + prewidth/2;
				const prewY  = baseInfo.previewCanvasWidth* y/baseInfo.paintingCanvasWidth + prewidth/2;
				
				
				// circleCtx.clearRect(0, 0, circleCanvas.height, circleCanvas.width);
				circleCanvas.height = circleCanvas.height
				// circleCtx.fillRect(prewX, prewY,prewidth,prewidth)
				circleCtx.beginPath();
				
				
				circleCtx.globalAlpha = alpha;
				
				
			
				
				
				circleCtx.arc(prewX, prewY,prewidth/2,0,2*Math.PI);
				circleCtx.fillStyle = '#FFFFFF';
				circleCtx.fill()
				
				circleCtx.arc(prewX, prewY,prewidth/3,0,2*Math.PI);
				circleCtx.fillStyle = color;
				circleCtx.fill()
				
				
				if (canvasX > 0 && canvasX < paintingCanvasWidth && canvasY > 0 && canvasY <
					paintingCanvasWidth && this.isDrawing) {

					
					this.mousePos = [x, y]
					
					
					
				}
			})




			paintingCanvas.addEventListener('mouseup', (e) => {
				this.isDrawing = false;

			})

		},
		initLayers() {
			let layersArray = [];
			let flod = false;
			for (let i = 0; i < 5; i++) {
				const item = {
					name: `Layer ${i + 1}`,
					id: `${i + 1}`,
					flod: flod,
					isChosen: false,
					images: [],
					// writable: false
				}
				let len = Math.floor(Math.random() * 5 + 1);
				let visible = true;
				for (let j = 0; j < len; j++) {
					const subitem = {
						name: `Image ${i + 1}-${j + 1}`,
						id: `${i + 1}-${j + 1}`,
						visible: visible,
						isChosen: false
					}
					item.images.push(subitem);
					visible = !visible;
				}
				layersArray.push(item)
				flod = !flod;
			}
			this.layersArray = layersArray;
		},
		draw(canvasX, canvasY) {

			if (this.isDrawing) {
				const baseInfo = this.baseInfo;
				const width = baseInfo.paintingCanvasWidth / baseInfo.paintingWidth; // 816/24




				let paintingCanvas = document.getElementById("painting-canvas");
				let paintingCtx = paintingCanvas.getContext("2d");
				paintingCtx.fillStyle = this.color;
				paintingCtx.globalAlpha = this.alpha;

				// 
				// editImageId 对应的位置 1 擦除

				let previewCanvas = document.getElementById("preview-canvas");
				let previewCtx = previewCanvas.getContext("2d");
				previewCtx.fillStyle = this.color;
				previewCtx.globalAlpha = this.alpha;

				let x = parseInt(canvasX / width);
				let y = parseInt(canvasY / width);
				let index = x + y * baseInfo.paintingWidth;

				const previewWidth = baseInfo.previewCanvasWidth / baseInfo.paintingWidth

				if (this.currentTool == 'eraser') {


					const layersArray = this.layersArray;
					const editImageId = this.editImageId;
					for (let i = 0; i < layersArray.length; i++) {
						const layer = layersArray[i];
						const images = layer.images;
						for (let j = 0; j < images.length; j++) {

							if (images[j].id == editImageId) {

								const colors = images[j].colors;



								for (let k = 0; k < colors.length; k++) {
									if (colors[k].pos[index]) {
										paintingCtx.clearRect(x * width, y * width, width, width);
										previewCtx.clearRect(parseInt(canvasX * baseInfo.previewCanvasWidth /
												baseInfo.paintingCanvasWidth / previewWidth) *
											previewWidth, parseInt(canvasY * baseInfo
												.previewCanvasWidth / baseInfo.paintingCanvasWidth /
												previewWidth) * previewWidth, previewWidth, previewWidth)
									}
								}

							}
						}
					}
				}


				if (this.currentTool == 'pencil') {
					paintingCtx.fillRect(x * width, y * width, width, width);

					previewCtx.fillRect(parseInt(canvasX * baseInfo.previewCanvasWidth / baseInfo
							.paintingCanvasWidth / previewWidth) * previewWidth, parseInt(canvasY *
							baseInfo.previewCanvasWidth / baseInfo.paintingCanvasWidth / previewWidth) *
						previewWidth, previewWidth, previewWidth)
				}

				//存储数据 
				let drawInfo = this.drawInfo;
				let pos = drawInfo.pos;

				//获取位置
				if (pos.length === 0) {
					pos = [];
					for (let n = 0; n < baseInfo.paintingWidth * baseInfo.paintingWidth; n++) pos[n] = 0
				}

				pos.splice(index, 1, 1);

				drawInfo.pos = pos;
				if (this.currentTool == 'pencil') {
					drawInfo.color = this.color;
					drawInfo.alpha = this.alpha;
				}
				this.drawInfo = drawInfo;

			}

		},
		//选择图层
		chooseLayer(i) {
			const layersArray = this.layersArray;
			const layer = layersArray[i];
			const editLayerId = layer.id;
			this.editLayerId = editLayerId;
			layer.flod = false;
			this.layersArray.splice(i, 1, layer);
			this.editImageId = '';
			this.setFrameInfo();
		},
		//显示图层下面的图片 只是展开折叠
		showImages(i) {
			let layersArray = this.layersArray;


			let layer = layersArray[i];
			layer.flod = !layer.flod;
			this.layersArray.splice(i, 1, layer);
		},
		//显示单个图片
		showImage(i, j) {
			let layersArray = this.layersArray;
			let layer = layersArray[i];

			let images = layer.images;
			for (let k = 0; k < images[k]; k++) {
				images[k].visible = false;
			}

			images[j].visible = true;

			this.layersArray.splice(i, 1, layer);

			// 


		},
		//增加一个图片
		addImage(i, type) {
			let layersArray = this.layersArray;
			let layer = layersArray[i];
			let images = layer.images;
			let count = this.addCounter('images', 'Image ' + layer.id);
			count++;
			this.newImageId++;
			images.push({
				name: `Image ${this.newImageId}`,
				id: `Image${this.newImageId}`, //`${layer.id}-${count}`,//`${i + 1}-${images.length + 1}`,
				visible: false,
				isChosen: false,
				colors: [], //提交的数据
				imageData: []
			})
			this.layersArray.splice(i, 1, layer);
			
			this.chooseLayer(i);
			this.setFrameInfo();
		},
		// 复制一个图层
		copyLayer(i) {

			let layersArray = this.layersArray;
			let layer = JSON.parse(JSON.stringify(layersArray[i]));

			const count = this.copyCounter('layers', layer.name);
			this.newLayerId++;
			let newLayer = {
				name: `${layer.name}(${count})`,
				id: `${this.newLayerId}`, //`${layer.id}-${count}`,//`${layersArray.length + 1}`,
				flod: true,
				isChosen: false,
				images: [],
			}
			
			this.layersArray.splice(i + count, 0, newLayer);
			this.afterCopyLayer(i,this.layersArray.length - 1)

		},
		afterCopyLayer(i, newIndex) {
			let layersArray = this.layersArray;
			let layer = layersArray[i];
		
			let newLayer = JSON.parse(JSON.stringify(layersArray[newIndex]))
			let paintingInfo = this.paintingInfo;
			const images = layer.images;
			for (let j = 0; j < images.length; j++) {
				const image = images[j]

				this.newImageId++;
				const id = `Image${this.newImageId}`
				let newImage = {
					name: `${image.name}`,
					id: id, //`${image.id} (${count + i})`,//`${ cpoyI + 1 } -${ layer.images.length + 1 } `,
					visible: false,
					isChosen: false
				}
				if (paintingInfo[image.id]) {
					paintingInfo[id] = JSON.parse(JSON.stringify(paintingInfo[image.id]))
				}

				newLayer.images.push(newImage);
			}

			this.layersArray.splice(newIndex, 1, newLayer)

			this.paintingInfo = paintingInfo;
			this.setFrameInfo();
		},
		//删除一个图层
		deleteLayer(i) {
			const layersArray = this.layersArray;
			const layer = layersArray[i];
			let paintingInfo = this.paintingInfo;
			const images = layer.images;
			for (let j = 0; j < images.length; j++) {
				delete paintingInfo[images[j].id];
			}
			this.layersArray.splice(i, 1);
			this.paintingInfo = paintingInfo;
			this.setPerviewImage();

			this.setEmptyCanvas();
			this.setFrameInfo();
		},
		// 增加一个图层
		addLayer() {
			let type = arguments[0]
			let layersArray = this.layersArray;
			let layer = layersArray[layersArray.length - 1];
			let count = this.addCounter('layers', 'Layer ');
			count++;
			this.newLayerId++;
			let newLayer = {
				name: `Layer ${count}`,
				id: `${this.newLayerId}`,
				flod: type ? false : true,
				isChosen: false,
				images: []
			}
			this.layersArray.splice(layersArray.length, 1, newLayer);

			// document.getElementById('layersBottom').scrollIntoView()
			// scrollTop()
			if (arguments[0] == 'init') {
				// this.chooseLayer(0)
				this.addImage(0, 'init')
			} else {
				this.setFrameInfo()
			}
			
		},
		toggleWritable(i) {
			let layersArray = this.layersArray;
			let layer = layersArray[i];
			// layer.writable = !layer.writable;
			// this.layersArray.splice(i, 1, layer);
			let j = arguments[1];
			if (typeof j == 'number') {
				const image = layer.images[j];
				this.editInfo = {
					name: '',
					placeholder: image.name,
					isEdit: 'image',
					i: i,
					j: j
				}
			} else {
				this.editInfo = {
					name: '',
					placeholder: layer.name,
					isEdit: 'layer',
					i: i
				}
			}
			
			$('.rename_modal').addClass('fadeIn').removeClass('fadeOut');
			
		},
		//重命名
		rename() {
			
			$('.rename_modal').addClass('fadeOut').removeClass('fadeIn');
			const editInfo = this.editInfo;

			if (editInfo.isEdit === 'layer') {
				// 图层
				let layersArray = this.layersArray;
				const i = editInfo.i;
				let layer = layersArray[i];

				layer.name = editInfo.name ? editInfo.name : editInfo.placeholder;
				this.layersArray.splice(i, 1, layer);
			} else {
				let layersArray = this.layersArray;
				const i = editInfo.i;
				let layer = layersArray[i];
				const j = editInfo.j;
				layer.images[j].name = editInfo.name ? editInfo.name : editInfo.placeholder;
				this.layersArray.splice(i, 1, layer);
			}

			this.editInfo = {
				name: '',
				placeholder: '',
				isEdit: '',
			}
			this.setFrameInfo();
		},
		//删除图片
		deleteImage(i, j) {
			const layersArray = this.layersArray;
			const layer = layersArray[i];
			const id = layer.images[j].id
			layer.images.splice(j, 1);
			this.layersArray.splice(i, 1, layer);
			let paintingInfo = this.paintingInfo;
			delete paintingInfo[id];
			this.paintingInfo = paintingInfo;
			this.setPerviewImage()
			this.setEmptyCanvas();
			this.setFrameInfo();

		},
		toggleCopyToast(i, j) {
			this.copyInfo = {
				isEdit: true,
				origin: 1,
				change: 1,
				i: i,
				j: j
			}
			$('.copy-img_modal').addClass('fadeIn').removeClass('fadeOut');
			
		},
		// 复制图片
		// 1.弹框选择张数 2.复制 
		copyImage() {
			$('.copy-img_modal').addClass('fadeOut').removeClass('fadeIn');
			const copyInfo = this.copyInfo;
			const cpoyI = copyInfo.i;
			const copyJ = copyInfo.j;
			let layersArray = this.layersArray;
			let layer = layersArray[cpoyI];
			let image = layer.images[copyJ];

			let paintingInfo = this.paintingInfo;
			const count = this.copyCounter('images', image.name);
			// 复制原图

			for (let i = 0; i < copyInfo.origin / 1; i++) {
				this.newImageId++;
				const id = `Image${this.newImageId}`
				let newImage = {
					name: `${image.name}(${count + i})`,
					id: id, //`${image.id} (${count + i})`,//`${ cpoyI + 1 } -${ layer.images.length + 1 } `,
					visible: false,
					isChosen: false
				}

				layer.images.splice(copyJ + 1 + i, 0, newImage);
				if (paintingInfo[image.id]) {
					paintingInfo[id] = JSON.parse(JSON.stringify(paintingInfo[image.id]))
				}




			}



			// 复制变色图片
			// 1.取到原图rgb 2.交换rgb位置获取新色

			// for (let i = 0; i < copyInfo.change / 1; i++) {
			//   this.newImageId++;
			//   const id = `Image${this.newImageId}`
			//   let newImage = {
			//     name: `${image.name}(${count + i + copyInfo.origin / 1})`,
			//     id: id,
			//     visible: false,
			//     isChosen: false
			//   }
			//   layer.images.splice(copyJ + i + 1 + copyInfo.origin / 1, 0, newImage);
			//   if (paintingInfo[image.id]) {

			//     const copyPaintingInfo = JSON.parse(JSON.stringify(paintingInfo[image.id]));
			//     for (let j = 0; j < copyPaintingInfo.length; j++) {
			//       let color = copyPaintingInfo[j].color;
			//       let r = color.slice(1, 3)
			//       let g = color.slice(3, 5)
			//       let b = color.slice(5, 7)
			//       const colorArray = [];
			//       colorArray[0] = "#" + r + b + g;
			//       colorArray[1] = "#" + b + r + g;
			//       colorArray[2] = "#" + g + r + b;
			//       console.log(i % 3)
			//       console.log(colorArray[i])
			//       copyPaintingInfo[j].color = colorArray[i]
			//     }

			//     paintingInfo[id] = copyPaintingInfo;
			//   }

			// }


			this.paintingInfo = paintingInfo;


			this.layersArray.splice(cpoyI, 1, layer);
			this.copyInfo = {
				isEdit: false,
				origin: 1,
				change: 1
			}
			
			this.setFrameInfo();
		},
		//移动图片即修改了图片层级 和显示隐藏图片相同 重新获取画布数据
		//选中编辑图片
		// 存储选中图片的Id 渲染的时候判断即可
		chooseImage(i, j) {
			const layersArray = this.layersArray;
			let layer = layersArray[i];
			console.log(layer.id)
			this.chooseLayer(i)
			let images = layer.images;

			for (let k = 0; k < images.length; k++) {
				images[k].visible = false;
			}
			let image = images[j]
			image.visible = true;

			const editImageId = images[j].id;
			this.editImageId = editImageId;

			layer.images.splice(j, 1, image)
			this.layersArray.splice(i, 1, layer)

			//画布上只显示当前选中图片的信息
			this.setFrameInfo();


		},
		//获取复制的名称
		copyCounter(type, name) {
			let count = 0;
			let nameObj = this.nameObj;
			if (type == 'layers') {
				const layers = nameObj.layers;
				let max = 0;
				for (let i = 0; i < layers.length; i++) {
					if (layers[i].indexOf(name) > -1) {

						count++;

						if (layers[i].split('(').length > 1 && layers[i].split('(')[layers[i].split('(')
								.length - 1].split(')')[0] / 1 > max) {
							max = layers[i].split('(')[layers[i].split('(').length - 1].split(')')[0] / 1
						}

					}
				}
				if (max == count)
					count++

			} else {
				const images = nameObj.images;
				let max = 0;
				for (let j = 0; j < images.length; j++) {
					if (images[j].indexOf(name) > -1) {
						count++;

						if (images[j].split('(').length > 1 && images[j].split('(')[images[j].split('(')
								.length - 1].split(')')[0] / 1 > max) {
							max = images[j].split('(')[images[j].split('(').length - 1].split(')')[0] / 1

						}
					}
				}

				if (max == count)
					count++
			}

			return count;
		},
		// 获取新增的名称
		addCounter(type, name) {
			let count = 0;
			let nameObj = this.nameObj;
			if (type == 'layers') {
				let layers = nameObj.layers;
				for (let i = 0; i < layers.length; i++) {
					if (layers[i].indexOf(name) > -1) {
						count++;
					}
				}
			} else {
				let images = nameObj.images;
				for (let j = 0; j < images.length; j++) {
					if (images[j].indexOf(name) > -1) {
						// count = images[j].split(name)[1] / 1 + 1;
						// count = images[j].split('-')[1] / 1 + 1;
						count++
					}
				}
			}

			return count;
		},
		//预览图片
		imagePreview() {
			let image = new Image();
			var canvas = document.getElementById("scale-canvas");
			//获取到预览图片的URL
			const paintingInfo = this.paintingInfo;
			if (paintingInfo == {}) {
				return
			}
			this.previewImage = canvas.toDataURL();
			this.isShow = true
			$('.preview-toast').addClass('fadeIn').removeClass('fadeOut')
		},
		
		hidePreview(){
			this.isShow = false
			$('.preview-toast').addClass('fadeOut').removeClass('fadeIn')
		},



		setPerviewImage() {
			let colorArray = [];
			let posArray = [];
			let alphaArray = [];
			const baseInfo = this.baseInfo;
			let previewCanvas = document.getElementById('preview-canvas');
			let previewCtx = previewCanvas.getContext('2d');



			for (let n = 0; n < baseInfo.paintingWidth * baseInfo.paintingWidth; n++) {
				posArray[n] = 0
				colorArray[n] = '';
				alphaArray[n] = 1;
			}

			previewCtx.clearRect(0, 0, baseInfo.previewCanvasWidth, baseInfo.previewCanvasWidth)



			let zoomIndexInfo = JSON.parse(JSON.stringify(this.zoomIndexInfo));
			zoomIndexInfo.reverse();
			let paintingInfo = this.paintingInfo;
			if (paintingInfo) {
				for (let i = 0; i < zoomIndexInfo.length; i++) {
					const imageid = zoomIndexInfo[i]
					if (paintingInfo[imageid]) {
						for (let k = 0; k < paintingInfo[imageid].length; k++) {
							let alpha = paintingInfo[imageid][k].alpha;
							let color = paintingInfo[imageid][k].color;
							let pos = paintingInfo[imageid][k].pos;
							for (let j = 0; j < pos.length; j++) {
								if (pos[j]) {
									posArray.splice(j, 1, pos[j])
									colorArray.splice(j, 1, color);
									alphaArray.splice(j, 1, alpha);

								}
							}
						}
					}


				}
			}

			let scaleCanvas = document.getElementById('scale-canvas')
			let scaleCtx = scaleCanvas.getContext('2d');


			const width = baseInfo.previewCanvasWidth / baseInfo.paintingWidth; // 816/24
			const scaleWidth = baseInfo.scaleCanvasWidth / baseInfo.paintingWidth;
			let isWhite = false;
			for (let i = 0; i < baseInfo.paintingWidth; i++) {
				isWhite = !isWhite;
				for (let j = 0; j < baseInfo.paintingWidth; j++) {
					isWhite = !isWhite;
					scaleCtx.globalAlpha = 0.8;
					scaleCtx.fillStyle = isWhite ? "#A1A8A1" : "#FFFFFF";
					scaleCtx.clearRect(scaleWidth * i, scaleWidth * j, scaleWidth, scaleWidth);
					scaleCtx.fillRect(scaleWidth * i, scaleWidth * j, scaleWidth, scaleWidth);
				}
			}


			for (let i = 0; i < baseInfo.paintingWidth; i++) {
				for (let j = 0; j < baseInfo.paintingWidth; j++) {
					if (posArray[i + j * baseInfo.paintingWidth]) {
						previewCtx.fillStyle = colorArray[i + j * baseInfo.paintingWidth];
						previewCtx.globalAlpha = alphaArray[i + j * baseInfo.paintingWidth];
						previewCtx.clearRect(width * i, width * j, width, width)
						previewCtx.fillRect(width * i, width * j, width, width)

						scaleCtx.fillStyle = colorArray[i + j * baseInfo.paintingWidth];
						scaleCtx.globalAlpha = alphaArray[i + j * baseInfo.paintingWidth];
						scaleCtx.clearRect(scaleWidth * i, scaleWidth * j, scaleWidth, scaleWidth)
						scaleCtx.fillRect(scaleWidth * i, scaleWidth * j, scaleWidth, scaleWidth)
					}

				}
			}




		},
		downloadFile() {
			const submitData = this.getSubmitData();
			if (!submitData) {
				
				showModal({
					message:'Please create works first!',
					type: 'error'
				})
				return
			}

			let txtStr = "";

			txtStr += JSON.stringify(submitData);
			let date = new Date();
			this.download("data" + date.getTime() + ".txt", txtStr);
		},

		download(filename, text) {
			var pom = document.createElement("a");
			pom.setAttribute(
				"href",
				"data:text/plain;charset=utf-8," + encodeURIComponent(text)
			);
			pom.setAttribute("download", filename);
			if (document.createEvent) {
				var event = document.createEvent("MouseEvents");
				event.initEvent("click", true, true);
				pom.dispatchEvent(event);
			} else {
				pom.click();
			}
		},
		getLayerName(layerId) {
			let layersArray = this.layersArray;
			let layerName = '';
			for (let i = 0; i < layersArray.length; i++) {
				if (layersArray[i].id == layerId) {
					layerName = layersArray[i].name
				}
			}
			return layerName;
		},
		getNewAlpha(alpha) {
			let x = alpha * 255;
			let y = Math.floor(x);
			let s = y.toString(16);
			let ret = y < 16 ? '0' + s : s;
			ret.toUpperCase();
			return ret;
		},

		getSubmitData() {
			let paintingInfo = this.paintingInfo;
			if (paintingInfo == {}) {

				return
			}
			const layersArray = this.layersArray;
			let groupName = {}
			for (let k in paintingInfo) {
				for (let i = 0; i < layersArray.length; i++) {
					const images = layersArray[i].images;
					// const layerName = layersArray[i].name;
					const layerId = layersArray[i].id;
					if (!groupName[layerId]) {
						groupName[layerId] = []
					}
					for (let j = 0; j < images.length; j++) {
						if (images[j].id == k) {
							groupName[layerId].push({
								name: images[j].name,
								id: images[j].id
							})
						}
					}
				}
			}
			let groups = []

			for (let key in groupName) {

				let layers = [];
				for (let i = 0; i < groupName[key].length; i++) {

					let name = groupName[key][i].name;
					let id = groupName[key][i].id;
					const imagesColor = paintingInfo[id];
					let data = {}
					for (let j = 0; j < imagesColor.length; j++) {
						const color = imagesColor[j].color.slice(1, 7)
						const alpha = this.getNewAlpha(imagesColor[j].alpha);
						let datakey = color + alpha;
						if (!data[datakey]) {
							data[datakey] = []
						}
						const pos = imagesColor[j].pos;
						for (let k = 0; k < pos.length; k++) {
							if (pos[k]) {
								data[datakey].push(k)
							}
						}
					}
					//删除空数组
					for (let color in data) {
						if (data[color].length == 0) {
							delete data[color]
						}
					}


					let suboption = {
						name: name,
						data: data
					}

					layers.push(suboption);

				}



				let layerName = this.getLayerName(key)
				let option = {
					name: layerName,
					layers: layers
				}
				groups.push(option)
			}
			if (groups.length == 0) {

				return
			}
			let date = new Date();

			const submitData = {
				groups: groups,
				name: date.getTime() + '',
				width: this.baseInfo.paintingWidth,
				height: this.baseInfo.paintingWidth,
				multiple: this.multiple
			}

			return submitData;
		},
		readTxt() {
			this.$refs.txt.click()
		},
		uploadFile(e) {
			//if(file)不存在 return

			const file = e.target.files[0];
			if (!file) return
			let name = file.name.split(".").splice(-1).toString();
			if (name !== "txt") {
				// this.$message({
				//   message: "文件类型错误,请重新选择文件",
				//   type: "warning",
				// });
				// this.infoToggle = true;
				// this.toastInfo = 'Invalid File Format<br/>(support for .TXT)'
				
				showModal({
					message:'Invalid File Format<br/>(support for .TXT)',
					type: 'error'
				})
				return;
			}
			let reader = new FileReader();
			if (typeof FileReader === "undefined") {
				// this.$message({
				//   message: "您的浏览器不支持",
				//   type: "warning",
				// });

				// this.infoToggle = true;
				// this.toastInfo = '您的浏览器不支持!'
				
				showModal({
					message:'您的浏览器不支持!',
					type: 'error',
				})

			}

			reader.onload = (e) => {
				let submitData = null;
				try {
					submitData = JSON.parse(reader.result)

					// this.submitData = JSON.parse(reader.result);
				} catch (error) {
					// this.$message({
					//   message: "文件数据格式错误,请重新选择文件",
					//   type: "warning",
					// });

					// this.infoToggle = true;
					// this.toastInfo = '文件数据格式错误,请重新选择文件!'
					
					showModal({
						message:'Invalid File Format<br/>(support for .TXT)',
						type: 'info',
						duration:5000
					})
					return;
				}

				console.log(submitData)
				if (submitData.groups.length > 0) {
					this.newLayerId = 0;
					this.newImageId = 0;
					let paintingInfo = {};

					let layersArray = [];

					const groups = submitData.groups;
					const width = submitData.width;
					const height = submitData.height;
					const multiple = submitData.multiple;
					this.multiple = multiple;
					
					for (let i = 0; i < groups.length; i++) {
						let group = groups[i];
						let layers = group.layers;
						let images = [];
						for (let j = 0; j < layers.length; j++) {
							let layer = layers[j];
							this.newImageId++;
							const imageId = this.newImageId;

							let imageOption = {
								name: layer.name,
								id: 'Image ' + imageId,
								visible: false,
								isChosen: false
							}
							images.push(imageOption)

							const data = layer.data;
							if (!paintingInfo[layer.name]) paintingInfo[layer.name] = []
							for (let colorinfo in data) {

								let color = '#' + colorinfo.slice(0, 6);
								let alpha = parseFloat(parseInt((parseInt(colorinfo.slice(6, 8), 16) /
									255) * 1000) / 1000).toFixed(2);
								let pos = [];
								let datapos = data[colorinfo]
								for (let n = 0; n < height * width; n++) pos[n] = 0;

								for (let point = 0; point < datapos.length; point++) {
									const colorIndex = datapos[point]
									pos[colorIndex] = 1
								}

								let paintingOption = {
									color: color,
									alpha: alpha,
									pos: pos
								}

								paintingInfo[layer.name].push(paintingOption)
							}


						}
						const id = this.newLayerId++;
						let layerOption = {
							name: group.name,
							id: 'Layer ' + id,
							flod: false,
							isChosen: false,
							images: images
						}
						layersArray.push(layerOption)

					}

					this.layersArray = layersArray;
					this.paintingInfo = paintingInfo;

					this.frameInfo = {
						layersArray: [],
						canvasInfo: [],
						imageArray: []
					}
					this.currentIndex = -1;
					this.setEmptyCanvas()
				}

			}

			reader.readAsText(file, "utf-8");

		},
		complete() {

			const submitData = this.getSubmitData();
			if (!submitData) {
				// this.infoToggle = true;
				// this.toastInfo = 'Please create works first.'
				showModal({
					message:'Please create works first!',
					type: 'info',
					duration:5000
				})
				return
			}
			showToast({
				type: 'loading',
				message: 'loading'
			})
			
			this.initWebSocket()

		},

		initWebSocket() {

			const ws = 'wss://pixel.coltstail.com/api/v1/draw-ws/121212'

			this.websocket = new WebSocket(ws)
			this.websocket.onopen = this.onOpenWebsocket
			this.websocket.onerror = this.onErrorWebsocket
			this.websocket.onclose = this.onCloseWebsocket
			this.websocket.onmessage = this.onMessageWebsocket
		},
		onOpenWebsocket() {
			const submitData = this.getSubmitData();
			let txtJson = JSON.stringify(submitData);

			this.websocket.send(txtJson)

		},
		onErrorWebsocket() {
			//TODO websocket 失败
			console.log("连接失败!");
		},
		onCloseWebsocket() {
			console.log("连接关闭!")
			hideToast();
		},
		onMessageWebsocket(message) {

			const res = JSON.parse(message.data)
			console.log(res)
			if (res.code == 0) {
				const data = res.data;
				if (data.finish) {
					const key = data.data;
					this.imageZipKey = key
					this.downloadImage()
				}
			}
		},

		downloadImage() {
			const key = this.imageZipKey;
			const requestUrl = `https://pixel.coltstail.com/api/v1/draw/${key}/download`
			console.log(requestUrl)
			$.ajax({
				url: requestUrl,
				method: 'get',
				dataType:'json',
				async: true,
				success:(res)=>{
					
					
					if(res.data){
						window.open(res.data);
						hideToast()
					} else {
						showToast({
							type: 'error',
							message: res.msg
						})
					}
					
				},
				error:(XmlHttpRequest,textStatus, errorThrown)=>{
					showToast({
						message:errorThrown,
						type: 'error',
						duration:5000
					})
				},
			})
			// request(requestUrl).then(res => {
			// 	console.log(res)
			// 	window.open(res)
			// 	this.loading = false
			// }).catch(err => {
			// 	// this.loading = false
			// 	// this.toastInfo = err;
			// 	showToast({
			// 		message:err,
			// 		type: 'error',
			// 		duration:5000
			// 	})
			// 	this.infoToggle = true;
			// })
		},
		changeFrame(type) {
			this.editImageId = '';
			console.log(this.editImageId)
			let index = this.currentIndex;
			console.log(index)
			if (type == '-') {
				index--
			} else {
				index++
			}
			this.currentIndex = index;
			const frameInfo = JSON.parse(JSON.stringify(this.frameInfo));
			if (index <= -1 || index > frameInfo.layersArray.length - 1) {
				return
			}
			
			
			


			console.log(index)
			const layersArray = JSON.parse(JSON.stringify(frameInfo.layersArray));
			this.layersArray = layersArray[index];
			const canvasInfo = JSON.parse(JSON.stringify(frameInfo.canvasInfo))
			this.paintingInfo = canvasInfo[index];
			
			console.log(this.paintingInfo)
			setTimeout(()=>{
				this.editImageId = frameInfo.imageArray[index];
			},10)
			
			
			
			console.log(this.editImageId)
			this.setPerviewImage();
			
		},

	}
})
