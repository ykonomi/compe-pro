## 概要

競プロで書いたコード、スニペット, 競技中にコードを書くファイルを保存する

## テクニック集

- 答えで二分探索
- 小さい制約は全探索を考えよう
- 木の直径は最短距離計算を２回やる
- 扱いやすい形にして前計算しよう
- 区間の総和は累積和
- 要素の探索はソートして二部探索
- パリティで考える
- 連携判定はunion-find
- 素数の逆数の和はO(loglogN)
- 主客転倒: ある点からの経路の合計を求めるではなく、辺をまたぐ経路がいくつあるかを考える
- 各頂点への最短経路はダイクストラ法
- 「定数倍」を見積もる
- 複数の部分問題に分解して、漸化式を立てることで答えを出す手法を動的計画法(DP)という
- 上界/下界を見積もる
- 領域加算は二次元いもす法
- (i & (1 << j)) = 0  => i の j ビット目（2^j の位）が 0 であるための条件。
- Ceil(x/n) = (x+1)/n,  Floor(x/n) = x/n
- int64 の最大値: 10^19 > 9_223_372_036_854_775_807 > 9.0 * 10^18
- aをbごとにわけると何グループできるか  (a + b - 1) / b
- 木: 連結で閉路を持たない無向グラフ, 木は二部グラフ
- 二部グラフ: 奇数長の閉路を含まない, 最大マッチングが多項式時間ですむ
