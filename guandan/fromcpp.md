// int GetCardsType(std::vector<int> cards) {
//     // TODO:
//     auto cnt = cards.size();

//     if (cnt == 1) {
//         return CTID_YI_ZHANG;
//     } else if (cnt == 2) {
//         return CTID_ER_ZHANG;
//     } else if (cnt == 3) {
//         return CTID_SAN_ZHANG;
//     }

//     std::map<int, int> colors;
//     std::map<int, int> numbs;
//     std::vector<int> jipaiCards;

//     for (auto& c : cards) {
//         if (c == m_wildCard) {
//             jipaiCards.emplace_back(c);
//         } else {
//             auto col = CCardFunc::cardcolor(c);
//             auto n = CCardFunc::cardnum(c);
//             colors[col]++;
//             numbs[n]++;
//         }
//     }

//     auto colorcnt = colors.size();
//     auto ncnt = numbs.size();
//     auto jcnt = jipaiCards.size();

//     FLOGD("{} {},{} {} {}", fmt::join(cards, "_"), cnt, colorcnt, ncnt, jcnt);

//     if (jcnt == 0) {
//         if (cnt == 4) {
//             if (ncnt == 1) {
//                 return CTID_SI_ZHANG;
//             } else if (ncnt == 2) {
//                 // CN_F;CN_Z;
//                 return CTID_SI_WANG;
//             }
//         } else if (cnt == 5) {
//             // 三带二
//             // 顺子
//             // 同花顺
//             // 5炸
//             if (ncnt == 5 && colorcnt == 1) {
//                 return CTID_TONG_HUA_SHUN;
//             } else if (ncnt == 2) {
//                 return CTID_SAN_DAI_ER;
//             } else if (ncnt == 5) {
//                 return CTID_SHUN_ZI;
//             } else if (ncnt == 1) {
//                 return CTID_WU_ZHANG;
//             }
//         } else if (cnt == 6) {
//             // 钢板, 6炸
//             if (ncnt == 2) {
//                 return CTID_GANG_BAN;
//             } else if (ncnt == 1) {
//                 return CTID_LIU_ZHANG;
//             } else if (ncnt == 3) {
//                 return CTID_SAN_LIAN_DUI;
//             }
//         } else if (cnt == 7) {
//             if (ncnt == 1) {
//                 return CTID_QI_ZHANG;
//             }
//         } else if (cnt == 8) {
//             if (ncnt == 1) {
//                 return CTID_BA_ZHANG;
//             }
//         }
//     } else if (jcnt == 1) {
//         if (cnt == 4) {
//             if (ncnt == 1) {
//                 return CTID_SI_ZHANG;
//             }
//         } else if (cnt == 5) {
//             // 三带二
//             // 顺子
//             // 同花顺
//             // 5炸
//             if (ncnt == 4 && colorcnt == 1) {
//                 return CTID_TONG_HUA_SHUN;
//             } else if (ncnt == 2) {
//                 return CTID_SAN_DAI_ER;
//             } else if (ncnt == 4) {
//                 return CTID_SHUN_ZI;
//             } else if (ncnt == 1) {
//                 return CTID_WU_ZHANG;
//             }
//         } else if (cnt == 6) {
//             // 钢板, 6炸
//             if (ncnt == 2) {
//                 return CTID_GANG_BAN;
//             } else if (ncnt == 1) {
//                 return CTID_LIU_ZHANG;
//             } else if (ncnt == 3) {
//                 return CTID_SAN_LIAN_DUI;
//             }
//         } else if (cnt == 7) {
//             if (ncnt == 1) {
//                 return CTID_QI_ZHANG;
//             }
//         } else if (cnt == 8) {
//             if (ncnt == 1) {
//                 return CTID_BA_ZHANG;
//             }
//         } else if (cnt == 9) {
//             if (ncnt == 1) {
//                 return CTID_JIU_ZHANG;
//             }
//         }
//     } else if (jcnt == 2) {
//         if (cnt == 4) {
//             if (ncnt == 1) {
//                 return CTID_SI_ZHANG;
//             }
//         } else if (cnt == 5) {

//             if (ncnt == 4 && colorcnt == 1) {
//                 return CTID_TONG_HUA_SHUN;
//             } else if (ncnt == 2) {
//                 return CTID_SAN_DAI_ER;
//             } else if (ncnt == 4) {
//                 return CTID_SHUN_ZI;
//             } else if (ncnt == 1) {
//                 return CTID_WU_ZHANG;
//             }
//         } else if (cnt == 6) {
//             if (ncnt == 2) {
//                 return CTID_GANG_BAN;
//             } else if (ncnt == 1) {
//                 return CTID_LIU_ZHANG;
//             } else if (ncnt == 3) {
//                 return CTID_SAN_LIAN_DUI;
//             }
//         } else if (cnt == 7) {
//             if (ncnt == 1) {
//                 return CTID_QI_ZHANG;
//             }
//         } else if (cnt == 8) {
//             if (ncnt == 1) {
//                 return CTID_BA_ZHANG;
//             }
//         } else if (cnt == 9) {
//             if (ncnt == 1) {
//                 return CTID_JIU_ZHANG;
//             }
//         } else if (cnt == 10) {
//             if (ncnt == 1) {
//                 return CTID_SHI_ZHANG;
//             }
//         }
//     } else {
//     }
//     return -1;
// }
// std::string Cards2StrWithRank(int type, std::vector<uint8_t> cards) {
//     if (cards.empty()) {
//         return "";
//     }

//     std::vector<uint8_t> normalCards;
//     std::vector<uint8_t> jipaiCards;

//     for (int i = 0; i < cards.size(); i++) {
//         if (cards[i] == m_wildCard) {
//             jipaiCards.emplace_back(cards[i]);
//         } else {
//             normalCards.emplace_back(cards[i]);
//         }
//     }

//     if (normalCards.empty()) {
//         normalCards.swap(jipaiCards);
//     }

//     std::string ret;
//     ret.reserve(cards.size() * 4);

//     if (!jipaiCards.empty()) {
//         // 此时要算集牌代表啥
//         switch (type) {
//         case CTID_YI_ZHANG: {
//             normalCards.emplace_back(jipaiCards[0]);
//         } break;
//         case CTID_ER_ZHANG: {
//             if (jipaiCards.size() == 2) {
//                 normalCards.swap(jipaiCards);
//             } else if (jipaiCards.size() == 1) {
//                 ret += CardWithRealRank(jipaiCards[0], CCardFunc::cardnum(normalCards[0])) + " ";
//             }
//         } break;
//         case CTID_SAN_DAI_ER: {
//             std::unordered_map<uint8_t, uint8_t> count;
//             std::for_each(normalCards.begin(), normalCards.end(), [&count](int num) { count[CCardFunc::cardnum(num)]++; });
//             // first, keep every type card cnt >= 2
//             for (auto& item : count) {
//                 while (!jipaiCards.empty()) {
//                     if (item.second < 2) {
//                         item.second++;
//                         ret += CardWithRealRank(jipaiCards[0], item.first) + " ";
//                         jipaiCards.pop_back();
//                     } else {
//                         break;
//                     }
//                 }
//             }
//             std::sort(normalCards.begin(), normalCards.end(), [this](const uint8_t& a, const uint8_t& b) {
//                 auto anum = CCardFunc::cardnum(a);
//                 auto bnum = CCardFunc::cardnum(b);
//                 if (anum == bnum) {
//                     return CCardFunc::cardcolor(a) < CCardFunc::cardcolor(b);
//                 }
//                 if (anum == m_wildCardNum) {
//                     return true;
//                 }
//                 if (bnum == m_wildCardNum) {
//                     return false;
//                 }
//                 return CCardFunc::cardnum(a) < CCardFunc::cardnum(b);
//             });

//             uint8_t maxNumCard = CN_NULL;
//             for (auto kitem : normalCards) {
//                 auto anum = CCardFunc::cardnum(kitem);
//                 if (anum == CN_B || anum == CN_F || anum == CN_Z || anum == CN_NULL) {
//                     continue;
//                 }
//                 maxNumCard = kitem;
//                 break;
//             }
//             while (!jipaiCards.empty()) {
//                 ret += CardWithRealRank(jipaiCards[0], CCardFunc::cardnum(maxNumCard)) + " ";
//                 jipaiCards.pop_back();
//             }
//         } break;
//         case CTID_SAN_LIAN_DUI: {
//             std::unordered_map<uint8_t, uint8_t> count;
//             std::for_each(normalCards.begin(), normalCards.end(), [&count](int num) { count[CCardFunc::cardnum(num)]++; });
//             if (count.size() == 3) {
//                 for (auto& item : count) {
//                     while (!jipaiCards.empty()) {
//                         if (item.second < 2) {
//                             item.second++;
//                             ret += CardWithRealRank(jipaiCards[0], item.first) + " ";
//                             jipaiCards.pop_back();
//                         } else {
//                             break;
//                         }
//                     }
//                 }
//             } else {
//                 auto maxNumCard = *std::max_element(normalCards.begin(), normalCards.end(), [](const uint8_t& a, const uint8_t& b) {
//                     return CCardFunc::cardnum(a) < CCardFunc::cardnum(b);
//                 });
//                 auto minNumCard = *std::min_element(normalCards.begin(), normalCards.end(), [](const uint8_t& a, const uint8_t& b) {
//                     return CCardFunc::cardnum(a) < CCardFunc::cardnum(b);
//                 });
//                 auto targetNum = 0;
//                 if (CCardFunc::cardnum(maxNumCard) < CN_A) {
//                     targetNum = CCardFunc::cardnum(maxNumCard) + 1;
//                 } else {
//                     targetNum = CCardFunc::cardnum(minNumCard) - 1;
//                 }
//                 while (!jipaiCards.empty()) {
//                     ret += CardWithRealRank(jipaiCards[0], targetNum) + " ";
//                     jipaiCards.pop_back();
//                 }
//             }
//         } break;
//         case CTID_SHUN_ZI:
//         case CTID_TONG_HUA_SHUN: {
//             // A2345; TJQKA; 23456
//             auto maxNumCard = *std::max_element(normalCards.begin(), normalCards.end(), [](const uint8_t& a, const uint8_t& b) {
//                 return CCardFunc::cardnum(a) < CCardFunc::cardnum(b);
//             });
//             auto minNumCard = *std::min_element(normalCards.begin(), normalCards.end(), [](const uint8_t& a, const uint8_t& b) {
//                 return CCardFunc::cardnum(a) < CCardFunc::cardnum(b);
//             });
//             auto distance = CCardFunc::cardnum(maxNumCard) - CCardFunc::cardnum(minNumCard) + 1;
//             auto expcnt = cards.size();

//             if (distance < expcnt) {
//                 // 2345,6; 235,46
//                 for (int startAt = CCardFunc::cardnum(minNumCard); startAt <= CN_A; startAt++) {
//                     if (jipaiCards.empty()) {
//                         break;
//                     }
//                     auto res = std::find_if(normalCards.begin(), normalCards.end(), [&startAt](const uint8_t& card) {
//                         return CCardFunc::cardnum(card) == startAt;
//                     });
//                     if (res == normalCards.end()) {
//                         ret += CardWithRealRank(jipaiCards[0], startAt) + " ";
//                         jipaiCards.pop_back();
//                     }
//                 }

//                 for (int startAt = CCardFunc::cardnum(minNumCard); startAt >= CN_2; startAt--) {
//                     if (jipaiCards.empty()) {
//                         break;
//                     }
//                     auto res = std::find_if(normalCards.begin(), normalCards.end(), [&startAt](const uint8_t& card) {
//                         return CCardFunc::cardnum(card) == startAt;
//                     });
//                     if (res == normalCards.end()) {
//                         ret += CardWithRealRank(jipaiCards[0], startAt) + " ";
//                         jipaiCards.pop_back();
//                     }
//                 }

//             } else {
//                 // A345,2; A245,3; A234,5;

//                 int startAt = CCardFunc::cardnum(minNumCard);
//                 int endAt = CCardFunc::cardnum(maxNumCard);

//                 if (endAt == CN_A && startAt != CN_10) {
//                     startAt = CN_2;
//                 }
//                 // only fill in
//                 for (; startAt < endAt; startAt++) {
//                     if (jipaiCards.empty()) {
//                         break;
//                     }
//                     auto res = std::find_if(normalCards.begin(), normalCards.end(), [&startAt](const uint8_t& card) {
//                         return CCardFunc::cardnum(card) == startAt;
//                     });
//                     if (res == normalCards.end()) {
//                         ret += CardWithRealRank(jipaiCards[0], startAt) + " ";
//                         jipaiCards.pop_back();
//                     }
//                 }
//             }
//         } break;
//         case CTID_GANG_BAN: {
//             std::unordered_map<uint8_t, uint8_t> count;
//             std::for_each(normalCards.begin(), normalCards.end(), [&count](int num) { count[CCardFunc::cardnum(num)]++; });
//             for (auto& item : count) {
//                 while (!jipaiCards.empty()) {
//                     if (item.second < 3) {
//                         item.second++;
//                         ret += CardWithRealRank(jipaiCards[0], item.first) + " ";
//                         jipaiCards.pop_back();
//                     } else {
//                         break;
//                     }
//                 }
//             }
//         } break;
//         case CTID_SAN_ZHANG:
//         case CTID_SI_ZHANG:
//         case CTID_WU_ZHANG:
//         case CTID_LIU_ZHANG:
//         case CTID_QI_ZHANG:
//         case CTID_BA_ZHANG:
//         case CTID_JIU_ZHANG:
//         case CTID_SHI_ZHANG: {
//             for (auto& card : jipaiCards) {
//                 ret += CardWithRealRank(card, CCardFunc::cardnum(normalCards[0])) + " ";
//             }
//         } break;
//         default:
//             break;
//         }
//     }

//     for (auto card : normalCards) {
//         ret += CardWithRealRank(card, CCardFunc::cardnum(card)) + " ";
//     }
//     if (ret.size() > 1 && ret.back() == ' ') {
//         ret.pop_back();
//     }
//     return ret;
// }