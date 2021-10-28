# Paquetes || Modulos || Versionamiento

[Regresar](https://gitlab.com/Alejo1595/go-desde-cero)

## Paquete

Es una carpeta que contiene una colección de archivos y nos provee una funcionalidad

## Modulos

Nos permite administar las dependencias de los paquetes y controlar el versionamiento. Pueden haber muchos modulos pero se recomienda que haya solo uno solo en la raiz de nustro proyecto.

Para crear un modulo, primero nos ubicamos en la raiz del proyecto y ejecutamos el siguiente comando:

```
go mod init nombre_modulo (Se recomienda que sea la misma ruta donde esta alojado el codigo fuente, es decir la ruta del repo)
```

Al ejecutar este comando se generara un archivo llamado go.mod el cual se usa para gestionar las dependencias de los paquetes.

Para llamar nuestros paquetes personalizados, debemos usar el nombre del modulo y le añadimos el nombre del paquete que queremos importar.

```
package main

import (
	"fmt"

	"github.com/Alejo1595/paquetes-GO/greet"
	"github.com/Alejo1595/paquetes-GO/test"
)

func main() {
	fmt.Println(greet.English())
	fmt.Println(test.Prueba())
}
```

Para importar paquetes de tercero, colocamos la ruta del paquete a utilizar y ejecutamos el siguiente comando:

```
go mod tidy
```

Este comando va a importar las dependencias que necesita el paquete de terceros que queremos utlizar.

Existen dos tipos de dependencias:

- Directas: Que son los paquetes de terceros que importamos
- Indirectas: Que son los paquetes que necesita los paquetes de terceros para funcionar.

Ademas de eso se va a generar un archivo llamado **go.sum** el cual va a llevar el control de las dependencias directas e indirectas del proyecto. Ademas este archivo tambien posee un hash por cada dependencia y este hash nos ayuda a validar que las dependencias no hayan sido modificadas.

## Versionamiento

Usando el siguiente comando se puede obtener información del modulo y las dependencias que este posee.

```
go list -m all
```

Este comando puede mosrtar las dependencias y la versión de cada una de ellas. GO se basa en el versionamiento de git para obtener el siguiente valor

```
v1.5.2

1 -> Mayor
5 -> Menor
2 -> Patch
```

y en caso de que la dependencia a usar no posea un versionamiento GO descargara el ultimo codigo fuente y creara un pseudo-versionamiento:

```
V0.0.0-fecha_ultimo-comit_hash-comit
```

El sistema de versionamiento se conoce como SemVer (Versionado Semántico) la cual tiene la siguiente estructura:

- La letra **v**
- Numero mayor
- Numero menor
- Numero patch
- Opcionar se puede puede añadir un indentificador para indicar que es una versión de pre-lanzamiento

Ejemplo

```
v.numero-Mayor.numero-Menor.numero-patch-identificador
v 2.5.3-beta
```

Cada cambio que se haga en el software se debe versionar, iniciando con el valor **v0.0.0** la cual indica que el desarrollo no deberia ser considerada estable. es decir que esta en desarrollo. Cuando el numero mayor cambia significa que el desarrollo ya es estable y se puede mostrar al publico **v1.0.0**.

- Numero patch -> Corrigue bugs y debe ser compatible con la versión anterior.
- Numero menor -> Cabios adicionales que no rompen la compatibilidad de las versiones anteriores del paquete dentro de la versión mayor. Cada vez que cambia este numero, el patch se debe reiniciar.
- Numero mayor -> Cambios que rompen la compatibilidad con las versiones anteriores del paquete.

Para saber el porque un paquete se importo, usamos el siguiente comando:

```
go mod why Nombre_paquete
```

Para entender el flujo de importación se debe leer de arriba hacia abajo.

Para conocer el listado de versiones de un paquete usamos el siguiente comando

```
go list -m -versions Nombre_Paquete
```

Cada vez que se haga un build, GO tratara de traer las ultimas versiones de los paquetes, pero cada uno de los paquetes que se necesitan para la aplicación puede especificar que version de otros paquetes necesita (por medio de su propio go.mod) y esta versión no necesariamente debe ser la ultima o incluso no tener version.

Los paquetes son guardatos en la siguiente ruta **GOPATH/pkg/mod**, para saber que ruta tiene el GOPATH usamos el siguiente comando:

```
go env GOPATH
```

Dentro de la carpeta de modulos podremos listar cada uno de los modulos utilizados y revisar internamente cada modulo para saber que versiones de otros paquetes requiere.

Para poder actualizar un paquete a su ultima versión usamos el siguiente comando:

```
go get Nombre_Paquete
```

Pero tambien se puede actualizar a una versión en especifico

```
go get Nombre_Paquete@vNumero_version
```

Para actualizar una version a una versión mayor GO propone generar un nuevo subdirectio "vNumero_nueva_version_mayor" en el cual este todas las nuevas caracteristicas de esa version, para poder actualizar a esa version, basta con colocar la ruta hacia el paquete mas el directorio que apunte a la version que se necesita.

```
rsc.io/quote -> rsc.io/quote/v3
```

Aqui en este punto pueden pasar dos cosas:

1. Se necesita trabajar con dos versiones del mismo paquete, en este caso necesitamos el uso de alias para poder indicarle a GO como diferenciar un paquete del otro

El alias se genera de la sigiente manera

```
alias ruta
```

Ejemplo: Se puede observar dos versiones del mismo paquete, uno con la versión 1.x.x y otro con la versión 3.x.x y para que ambos puedan subsistir juntos se usa un alias

```
"rsc.io/quote" -> Este apunta a la versión 1.x.x
quoteV3 () "rsc.io/quote/v3" -> Este apunta a la versión 3.x.x
```

2. Se necesita actualizar de una version antigua a una mas nueva, para ello se borra la importación que hace referencia a una versión antigua.

En ambos casos es recomendable usar el comando

```
go mod tidy
```

Para poder actualizar las dependencias indirectas.

[Regresar](https://gitlab.com/Alejo1595/go-desde-cero)