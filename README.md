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