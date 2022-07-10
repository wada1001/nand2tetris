# nand2tetris

## 01 boolean logic

Nandを使用して幾つかの命令を実装

・Mux(マルチプレクサ)  
　bit反転を使って有効な桁のみを出す  
　inが増えても同様
・ DMux  
　inをbit反転して下に下ろす  

## 02 boolean arithmetic

・b=false  
　0埋め  

・Mux16(a=d1, b=notd1, sel=no, out=out, out[15]=ng, out[0..7]=l, out[8..15]=h);  
　戻り値は複数とれる  

## 03 Sequential logic

・DFF  
　時間の変異とデータ入力から値を保存する  
　正直なところ理解できていない気がする  

・Mux(a=out1, b=in, sel=load, out=out2);  
　未定義変数は0が暗黙的に入る？  

## 04 Machine Language

・機械語 => ニーモニック => アセンブリ  
　より人間に読みやすくなる。アセンブリはアセンブラによって解釈される  

・アドレッシングモード  
　直接 => アドレスを指定して読み込み  
　即値 => 定数の読み込み  
　関節 => 配列などの高水準コマンドにより確保されるアドレス  
　　物理的にj離れた位置に確保する  
　　例) c[j] == *(c + j)  

・アドレス空間  
　マシンごとに別々のメモリレイアウトがある  
　それがわからないうちは手がつけられなかった  

## 05 Computer Architecture

・CPU  
　A,Dメモリの役割、 A,C C命令の動作が重要だった  
　全体として処理は今までより多くなったが、命令を砕いてスコープを狭めればわかりやすい。  
　ALUの戻り値のzr, ngの意味がわかった  
