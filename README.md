# ASCIIArt

Imprime cada letra formada con caracteres del texto se especifique como entrada.


### Ejemplo
```Go
fmt.Println(chars.Parse("hola"))
```
Input: Hola
Output: 
```
**    **    *******     **             **       
**    **  ***     ***   **            ****      
**    ** ***        *** **           **  **     
******** ***        *** **          **    **    
******** ***        *** **         **********   
**    **  ***     ***   ********  ************  
**    **    *******     ******** **          **
```

Tambien se puede asignar el caracter con el cual se quiere imprimir cada letra.
```Go
fmt.Println(chars.Parsef("hola", '♠'))
```
Ouput:
```
♠♠    ♠♠    ♠♠♠♠♠♠♠     ♠♠             ♠♠       
♠♠    ♠♠  ♠♠♠     ♠♠♠   ♠♠            ♠♠♠♠      
♠♠    ♠♠ ♠♠♠        ♠♠♠ ♠♠           ♠♠  ♠♠     
♠♠♠♠♠♠♠♠ ♠♠♠        ♠♠♠ ♠♠          ♠♠    ♠♠    
♠♠♠♠♠♠♠♠ ♠♠♠        ♠♠♠ ♠♠         ♠♠♠♠♠♠♠♠♠♠   
♠♠    ♠♠  ♠♠♠     ♠♠♠   ♠♠♠♠♠♠♠♠  ♠♠♠♠♠♠♠♠♠♠♠♠  
♠♠    ♠♠    ♠♠♠♠♠♠♠     ♠♠♠♠♠♠♠♠ ♠♠          ♠♠ 
```
