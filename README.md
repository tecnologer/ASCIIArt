# ASCIIArt

Imprime cada letra formada con caracteres del texto se especifique como entrada.


### Ejemplos
**Ejemplo 1:** Imprimir "Hola" con asteriscos (es el caracter default), se usa la funcion `Parse`. Recibe la cadena que se desea imprimir con asteriscos.
```Go
fmt.Println(chars.Parse("hola"))
```
*Input:* **Hola** <br>
*Output: *
```
**    **    *******     **             **       
**    **  ***     ***   **            ****      
**    ** ***        *** **           **  **     
******** ***        *** **          **    **    
******** ***        *** **         **********   
**    **  ***     ***   ********  ************  
**    **    *******     ******** **          **
```

**Ejemplo 2:** Imprimir "Hola" con un caracter que no sea un asterisco. Para eso se usa la funcion `Parsef`. Que recibe la cadena a transformar y el caracter con el cual se imprimiran las letras.
```Go
fmt.Println(chars.Parsef("hola", '♠'))
```
*Ouput:*
```
♠♠    ♠♠    ♠♠♠♠♠♠♠     ♠♠             ♠♠       
♠♠    ♠♠  ♠♠♠     ♠♠♠   ♠♠            ♠♠♠♠      
♠♠    ♠♠ ♠♠♠        ♠♠♠ ♠♠           ♠♠  ♠♠     
♠♠♠♠♠♠♠♠ ♠♠♠        ♠♠♠ ♠♠          ♠♠    ♠♠    
♠♠♠♠♠♠♠♠ ♠♠♠        ♠♠♠ ♠♠         ♠♠♠♠♠♠♠♠♠♠   
♠♠    ♠♠  ♠♠♠     ♠♠♠   ♠♠♠♠♠♠♠♠  ♠♠♠♠♠♠♠♠♠♠♠♠  
♠♠    ♠♠    ♠♠♠♠♠♠♠     ♠♠♠♠♠♠♠♠ ♠♠          ♠♠ 
```

### Instalacion
`go get github.com/tecnologer/ASCIIArt`

### Ejemplo en codigo
```Go
package main

import (
	"fmt"

	"github.com/tecnologer/Strings"
)

func main() {
	fmt.Println(chars.Parsef("hola", '♠'))
}
```
