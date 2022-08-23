package ginrummy

// c++ code:
// float CalCardImpoveScore(CardsPtr incards, CMyCard incard) {
// 	auto cards = incards->Clone();

// 	std::vector<CardsPtr> beforeMelds;
// 	CardsPtr beforeDeadwood = Cards::Create();
// 	GinRummy::DetectBest(cards, beforeMelds, beforeDeadwood);
// 	auto beforeScore = GinRummy::CardsPoint(beforeDeadwood);
// 	beforeScore = beforeScore == 0 ? 1 : beforeScore;

// 	cards->PushBack(incard);
// 	std::vector<CardsPtr> newMelds;
// 	CardsPtr newDeadwood = Cards::Create();
// 	GinRummy::DetectBest(cards, newMelds, newDeadwood);
// 	auto newScore = GinRummy::CardsPoint(newDeadwood);

// 	// 开始计算`holdrate`
// 	// 更新概率
// 	rate.InitRate(1.f / float(drawStackCount + 1));
// 	rate.SetCardsRate(discardstack, 0.f);
// 	rate.SetCardsRate(handCards, 0.f);
// 	rate.SetCardsRate(otherPlayerCards, 0.f);

// 	decltype(rate) newRate;
// 	rate.Clone(newRate);

// 	rate.SetCardsRate(beforeDeadwood, 1.0f);
// 	auto beforeRateList = rate.HoldpowerRateByOrder(beforeDeadwood);

// 	newRate.SetCardsRate(newDeadwood, 1.0f);
// 	auto newRateList = rate.HoldpowerRateByOrder(newDeadwood);

// 	// 计算下打掉那张牌
// 	auto discard = CMyCard(0);
// 	if (newRateList.size() > 0) {
// 		discard = newRateList[0]->card;
// 	}
// 	newScore = newScore - GinRummy::CardPoint(discard);

// 	float retscore = 0.f;
// 	std::string strlog;

// 	do {
// 		// 不要废牌
// 		if (incard.GetID() == discard.GetID()) {
// 			strlog = "c1-discard";
// 			retscore = 0.f;
// 			break;
// 		}
// 		auto meldChange = newMelds.size() - beforeMelds.size();
// 		// 能组成新的 meld.
// 		if (meldChange > 0) {
// 			retscore = float(meldChange * 10);
// 			strlog = fmt::format("c1-meldChange:{}", meldChange);
// 			break;
// 		}

// 		// 能融入先有的 meld
// 		if (beforeDeadwood->Size() >= newDeadwood->Size()) {
// 			retscore = float(GinRummy::CardPoint(incard));
// 			strlog = fmt::format("c2-merge_in_meld:{}", retscore);
// 			break;
// 		}

// 		// 降低的分数
// 		float improveScore = beforeScore - newScore;
// 		float improveScoreRate = float(improveScore) / float(beforeScore);

// 		RLOGI(
// 			fmt::format("CalCardImpoveScore card:{}, beforeScore:{}, newScore:{}, improveScore:{}, improveScoreRate:{}", discard.GetName(), beforeScore, newScore, improveScore, improveScoreRate));

// 		float beforeRateSum = 0.f;
// 		for (auto& item : beforeRateList) {
// 			beforeRateSum += item->sum;
// 		}

// 		float newRateSum = 0.f;
// 		for (int i = 1; i < newRateList.size(); i++) {
// 			newRateSum += newRateList[i]->sum;
// 		}
// 		const float fixRate = 1.0f;
// 		auto improveGotRateRate = float(newRateSum - beforeRateSum) / float(beforeRateSum);

// 		// 提高分数比, 提高概率比 -> retscore > 0 要
// 		// 降低了概率, 降低了分数 ->
// 		// 提高了概率, 但是增加了分数 ->?
// 		// 降低了分数也降低了概率 -> retscore <= 0 不要

// 		retscore = improveScoreRate + improveGotRateRate;
// 		strlog = fmt::format("c3-Score:{}, ScoreRate:{}, GotRate:{}, GotRateRate:{}", improveScore, improveScoreRate, newRateSum - beforeRateSum, improveGotRateRate);
// 	} while (false);

// 	RLOGI(fmt::format("cal-card: {} to {},score:{},discard:{},status:{}", incard.GetName(), incards->Chinese(), retscore, discard.GetName(), strlog));
// 	return retscore;
// }

// void DoHandoutCard() {
// 	if (handCards->Size() != 11) {
// 		RLOGW(fmt::format("DoHandoutCard, the cards size is {}, and cards: {}", handCards->Size(), handCards->Chinese()));
// 	}

// 	std::vector<CardsPtr> melds;
// 	CardsPtr deadwood = Cards::Create();
// 	GinRummy::DetectBest(handCards, melds, deadwood);

// 	unsigned char outCard = 0;
// 	if (deadwood->Size() == 0) {
// 		RLOGW("DoHandoutCard, cal deadwood size is 0");
// 	} else if (deadwood->Size() == 1) {
// 		outCard = deadwood->Get(0).GetID();
// 	} else {
// 		rate.InitRate(1.f / float(drawStackCount + 1));
// 		rate.SetCardsRate(discardstack, 0.f);
// 		rate.SetCardsRate(handCards, 0.f);
// 		rate.SetCardsRate(deadwood, 1.0f);

// 		auto newRateList = rate.HoldpowerRateByOrder(deadwood);
// 		auto& res0 = newRateList[0];
// 		auto& res1 = newRateList[1];

// 		outCard = res0->card.GetID();
// 		if (!CanHandoutCard(res0->card)) {
// 			RLOGW("discard is drawcard");
// 			outCard = res1->card.GetID();
// 		}
// 		RLOGI(fmt::format("DoHandoutCard card0:{}, card1:{}", res0->String(), res1->String()));
// 	}
// }
