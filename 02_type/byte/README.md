## byteタイプ

1バイトは8ビットです。2^8範囲のデータを保存できます。

1byteはハードウェアでデータを読み込み、書き込みする際の最小単位です。メモリで扱えるboolean以外の最小単位のデータ保存タイプを作る場合はbyte単位で十分で、もっと小さい単位のデータタイプは要りません。

ですから、ハードウェアに保存するデータを扱う際はbyte単位で扱うのが良く、画像、動画などを扱う際のバイナリデータを扱う場合使用するタイプです。