package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Manejar la ruta principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Petición %s recibida por Devopsaurios %s", r.Method, r.URL.Path)
		// Escribir un encabezado HTML básico
		fmt.Fprintf(w, `
        <!DOCTYPE html>
		<html lang="es">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title> Producto 1 </title>
			<style>
				body {
					font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
					margin: 0;
					padding: 0;
					background-color:rgb(245, 245, 245);
					color: #333;
					line-height: 1.6;
				}

				.container {
					max-width: 1000px;
					margin: 0 auto;
					padding: 20px;
					background-color: white;
					text-align: center;
					box-shadow: 0 0 10px rgba(0,0,0,0.1);
					border-radius: 8px;
					margin-top: 30px;
				}
				.content {
					text-align: center;
				}

				header {
					background-color: #73EDFF;
					color: #000;
					padding: 20px 0;
					text-align: center;
					border-radius: 8px 8px 0 0;
				}	

				.button {
					display: inline-block;
					background-color: #73EDFF;
					color: #000;
					padding: 12px 24px;
					text-decoration: none;
					border-radius: 30px;
					font-weight: bold;
					margin-top: 20px;
					transition: all 0.3s ease;
					border: 2px solid #000;
				}
				.button:hover {
					background-color: #000;
					color: #73EDFF;
					transform: translateY(-3px);
					box-shadow: 0 6px 12px rgba(0,0,0,0.2);
				}
				
				.footer {
					max-width: 1000px;
					margin: 0 auto;
					padding: 20px;
					background-color: white;
					background-color: #73EDFF;
					box-shadow: 0 0 10px rgba(0,0,0,0.1);
					border-radius: 8px;
					margin-top: 30px;
					text-align: center;
				}

			</style>
		</head>

		<body>
		<div class= "container">
			<header>
				<h1> Soy alumno de la UOC, y somos el grupo DevOpsaurios </h1>
			</header>
		
		<div class= "content">
				<h2> Bienvenidos a nuestra primera entrega del Producto 1 </h2>
				<img src="/imagen" alt="Imagen UOC" width="300">
				<br>
				<a href="https://sites.google.com/uoc.edu/devopsauriosoficial/miembros-del-equipo" class="button" target="_blank">DevOpsaurios</a>
		</div>

			<footer>
				<p> &copy; FP.050 - Devops y Cloud Computing - Tutor: Jose Antonio Ramirez Ramirez </p>
			</footer>

		</div>
	</body>
    </html>
        `)
	})

	// Manejar la ruta para la imagen
	http.HandleFunc("/imagen", func(w http.ResponseWriter, r *http.Request) {
		// Servir una imagen de ejemplo (reemplaza con tu propia imagen)
		http.ServeFile(w, r, "uoc-logo.jpg")
	})

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor iniciado en http://localhost:8080")
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
