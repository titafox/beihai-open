package models

type Line struct {
	Id   int    `json:"id"`
	Name string `json:"name" comment:"作品名称"`
	List []Info `json:"list"`
}

type Info struct {
	No    int     `json:"no" comment:"作品编号"`
	Time  int64   `json:"time" comment:"支付时间"`
	User  string  `json:"user" comment:"支付用户"`
	Money float64 `json:"money" comment:"金额"`
}

func GetLine() []Line {
	var line []Line
	err := Eloquent.Raw("select id,name from tbl_works where status=1").Scan(&line).Error
	if err != nil {
		return make([]Line, 0, 0)
	}

	for k, v := range line {
		sql := "select n.work_no as no ,o.pay_time as time,CONCAT(LEFT(a.phone,3),'****',RIGHT(a.phone,4)) as user,o.order_money as money from  tbl_orders o,tbl_account a,tbl_nfts n where  o.typ=2 and o.wid = ? and o.status =2 and o.account_id=a.id and o.nft_id = n.id"
		var info []Info
		err := Eloquent.Raw(sql, v.Id).Scan(&info).Error
		if err != nil {
			line[k].List = make([]Info, 0, 0)
		} else {
			line[k].List = info
		}
	}

	return line

}

type LineInfo struct {
	//HasNext bool `json:"has_next"`
	Series            []Serie                 `json:"series"`
	AreaSeries        []AreaSeriesInfo        `json:"areaSeries"`
	CandlestickSeries []CandlestickSeriesInfo `json:"candlestickSeries"`
	//List []DownList `json:"list"`
}

type AreaSeriesInfo struct {
	Time  string  `json:"time"`
	Value float64 `json:"value"`
}

type CandlestickSeriesInfo struct {
	Time  string  `json:"time"`
	Open  float64 `json:"open"`
	High  float64 `json:"high"`
	Low   float64 `json:"low"`
	Close float64 `json:"close"`
}

type DoneList struct {
	Phone string  `json:"phone"`
	Price float64 `json:"price"`
	Time  string  `json:"time"`
}

type LinePrice struct {
	Price float64 `json:"price"`
}

func GetLine2(wid int64) *LineInfo {
	//查询品牌
	var lineInfo LineInfo
	var series []Serie
	err := Eloquent.Raw("select id,name from tbl_works where status=1").Scan(&series).Error
	if err != nil {
		lineInfo.Series = make([]Serie, 0, 0)
	} else {
		lineInfo.Series = series
	}

	if wid == 0 {
		wid = 2
	}

	//AreaSeries
	sql := "select FROM_UNIXTIME(pay_time, '%Y-%m-%d') as time,format(sum(price) /count(*),2) as value from tbl_orders where wid= ? and nft_id !='' and status=2 group by FROM_UNIXTIME(pay_time, '%Y-%m-%d')"
	var areaSeries []AreaSeriesInfo
	err = Eloquent.Raw(sql, wid).Scan(&areaSeries).Error
	if err != nil && len(areaSeries) == 0 {
		lineInfo.AreaSeries = make([]AreaSeriesInfo, 0, 0)
		lineInfo.CandlestickSeries = make([]CandlestickSeriesInfo, 0, 0)
	}
	lineInfo.AreaSeries = areaSeries

	candlestickSeriesInfo := make([]CandlestickSeriesInfo, len(areaSeries), len(areaSeries))
	for k, v := range areaSeries {
		candlestickSeriesInfo[k].Time = v.Time
		//开盘价
		var linePrice LinePrice
		Eloquent.Raw("select price from tbl_orders where wid= ? and nft_id !='' and status=2 and FROM_UNIXTIME(pay_time, '%Y-%m-%d')=? order by id limit 1", wid, v.Time).Scan(&linePrice)
		candlestickSeriesInfo[k].Open = linePrice.Price

		//收盘价
		Eloquent.Raw("select price from tbl_orders where wid= ? and nft_id !='' and status=2 and FROM_UNIXTIME(pay_time, '%Y-%m-%d')=? order by id desc limit 1", wid, v.Time).Scan(&linePrice)
		candlestickSeriesInfo[k].Close = linePrice.Price

		//最高价
		Eloquent.Raw("select price from tbl_orders where wid= ? and nft_id !='' and status=2 and FROM_UNIXTIME(pay_time, '%Y-%m-%d')=? order by price desc limit 1", wid, v.Time).Scan(&linePrice)
		candlestickSeriesInfo[k].High = linePrice.Price

		//最低价
		Eloquent.Raw("select price from tbl_orders where wid= ? and nft_id !='' and status=2 and FROM_UNIXTIME(pay_time, '%Y-%m-%d')=? order by price limit 1", wid, v.Time).Scan(&linePrice)
		candlestickSeriesInfo[k].Low = linePrice.Price
	}
	lineInfo.CandlestickSeries = candlestickSeriesInfo

	//lineInfo.CandlestickSeries = []CandlestickSeriesInfo{
	//	{Time:"2022-11-14",Open:234.5,High: 256.7,Low: 222.7,Close: 231.5},
	//	{Time:"2022-11-15",Open:234.5,High: 256.7,Low: 222.7,Close: 231.5},
	//	{Time:"2022-11-16",Open:234.5,High: 256.7,Low: 222.7,Close: 231.5},
	//}

	//lineInfo.List = []DownList{
	//{Phone: "13800000001",Money: 342.4,Time: "2022-11-14 16:30:42"},
	//{Phone: "13800000002",Money: 347.4,Time: "2022-11-15 16:30:42"},
	//}

	return &lineInfo
}

func GetdownList(wid int64, date string) []DoneList {
	var list []DoneList
	err := Eloquent.Raw("select CONCAT(LEFT(a.phone,3),'****',RIGHT(a.phone,4)) as phone, o.price,FROM_UNIXTIME(o.pay_time, '%Y-%m-%d %H:%i:%s') as time from tbl_orders o,tbl_account a where o.account_id=a.id and o.wid= ? and o.nft_id !='' and o.status=2 and FROM_UNIXTIME(o.pay_time, '%Y-%m-%d')=? order by o.id desc", wid, date).Scan(&list).Error
	if err != nil && len(list) == 0 {
		return make([]DoneList, 0, 0)
	}
	return list
}
