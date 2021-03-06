package goctp

func NewSnapshot(pDepthMarketData CThostFtdcDepthMarketDataField) (*Snapshot, error) {
	snapshot := &Snapshot{}
	snapshot.TradingDay = string(pDepthMarketData.GetTradingDay())
	snapshot.InstrumentID = string(pDepthMarketData.GetInstrumentID())
	snapshot.ExchangeInstID = string(pDepthMarketData.GetExchangeInstID())
	snapshot.LastPrice = pDepthMarketData.GetLastPrice()
	snapshot.PreSettlementPrice = pDepthMarketData.GetPreSettlementPrice()
	snapshot.PreClosePrice = pDepthMarketData.GetPreClosePrice()
	snapshot.PreOpenInterest = pDepthMarketData.GetPreOpenInterest()
	snapshot.OpenPrice = pDepthMarketData.GetOpenPrice()
	snapshot.HighestPrice = pDepthMarketData.GetHighestPrice()
	snapshot.LowestPrice = pDepthMarketData.GetLowestPrice()
	snapshot.Volume = int64(pDepthMarketData.GetVolume())
	snapshot.Turnover = pDepthMarketData.GetTurnover()
	snapshot.OpenInterest = pDepthMarketData.GetOpenInterest()
	snapshot.ClosePrice = pDepthMarketData.GetClosePrice()
	snapshot.SettlementPrice = pDepthMarketData.GetSettlementPrice()
	snapshot.UpperLimitPrice = pDepthMarketData.GetUpperLimitPrice()
	snapshot.LowerLimitPrice = pDepthMarketData.GetLowerLimitPrice()
	snapshot.PreDelta = pDepthMarketData.GetPreDelta()
	snapshot.CurrDelta = pDepthMarketData.GetCurrDelta()
	snapshot.UpdateTime = pDepthMarketData.GetUpdateTime()
	snapshot.UpdateMillisec = int64(pDepthMarketData.GetUpdateMillisec())
	snapshot.BidPrice1 = pDepthMarketData.GetBidPrice1()
	snapshot.BidVolume1 = int64(pDepthMarketData.GetBidVolume1())
	snapshot.AskPrice1 = pDepthMarketData.GetAskPrice1()
	snapshot.AskVolume1 = int64(pDepthMarketData.GetAskVolume1())
	snapshot.BidPrice2 = pDepthMarketData.GetBidPrice2()
	snapshot.BidVolume2 = int64(pDepthMarketData.GetBidVolume2())
	snapshot.AskPrice2 = pDepthMarketData.GetAskPrice2()
	snapshot.AskVolume2 = int64(pDepthMarketData.GetAskVolume2())
	snapshot.BidPrice3 = pDepthMarketData.GetBidPrice3()
	snapshot.BidVolume3 = int64(pDepthMarketData.GetBidVolume3())
	snapshot.AskPrice3 = pDepthMarketData.GetAskPrice3()
	snapshot.AskVolume3 = int64(pDepthMarketData.GetAskVolume3())
	snapshot.BidPrice4 = pDepthMarketData.GetBidPrice4()
	snapshot.BidVolume4 = int64(pDepthMarketData.GetBidVolume4())
	snapshot.AskPrice4 = pDepthMarketData.GetAskPrice4()
	snapshot.AskVolume4 = int64(pDepthMarketData.GetAskVolume4())
	snapshot.BidPrice5 = pDepthMarketData.GetBidPrice5()
	snapshot.BidVolume5 = int64(pDepthMarketData.GetBidVolume5())
	snapshot.AskPrice5 = pDepthMarketData.GetAskPrice5()
	snapshot.AskVolume5 = int64(pDepthMarketData.GetAskVolume5())
	snapshot.AveragePrice = pDepthMarketData.GetAveragePrice()
	snapshot.ActionDay = pDepthMarketData.GetActionDay()

	return snapshot, nil
}
