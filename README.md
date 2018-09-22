# Load AverageでLEDの点滅を変化させる

* Raspberry Piの負荷が高くなると、GPIO21に接続したLEDの点滅が早くなります。
* LEDは300〜1KΩの抵抗を挟んで接続してください。


## 必要なもの

* go 1.11
* make

## ビルド

	$ make

## 実行

	$ ./blink

Verbose Mode

	$ ./blink -v

Help

	$ ./blink -h
