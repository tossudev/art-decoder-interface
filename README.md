░█▀█░█▀▄░▀█▀░░░█▀▄░█▀▀░█▀▀░█▀█░█▀▄░█▀▀░█▀▄  
░█▀█░█▀▄░░█░░░░█░█░█▀▀░█░░░█░█░█░█░█▀▀░█▀▄  
░▀░▀░▀░▀░░▀░░░░▀▀░░▀▀▀░▀▀▀░▀▀▀░▀▀░░▀▀▀░▀░▀  

A web tool which encodes art data into text-based (ASCII) art and decodes it back.

## 🍔 Prerequisites

Go (1.26<)  

## 🍕 Building & Usage

1. Clone the repository  
2. cd into the root of the cloned directory.  
3. Run ```go run .```  
4. Open ```localhost:6969``` in your browser.  
5. Insert art or encoded art in the text field and press corresponding buttons.  

## 🌮 Examples

### Basic decode
Input:  
```[5 #][5 -_]-[5 #]```  
Output:  
```#####-_-_-_-_-_-#####```  
  
### Lion art  
Input:  
```[8  ]@|\[2 @]
[7  ]-[2  ][4 @]
[6  ]/7[3  ][4 @]
[5  ]/[4  ][6 @]
[5  ]\-' [8 @]`-[15 _]
[6  ]-[9 @][13  ]/[4  ]\
 [7 _]/[4  ]/_[7  ][6 _]/[6  ]|[10 _]-
/,[10 _]/  `-.[3 _]/,[13 _][10 -]_)
```  
Output:  
```        @|\@@
       -  @@@@
      /7   @@@@
     /    @@@@@@
     \-' @@@@@@@@`-_______________
      -@@@@@@@@@             /    \
 _______/    /_       ______/      |__________-
/,__________/  `-.___/,_____________----------_)```  