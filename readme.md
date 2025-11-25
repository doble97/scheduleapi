# 🚀 miapigo: Clon Simple de Trello (API REST)

## 📝 Descripción del Proyecto

# ScheduleAPI — Clean Architecture

This repository follows a Clean Architecture layout for a Go service. The structure was created to separate concerns (domain, use cases, repositories, infrastructure, and interfaces) and to be easy to test.

Este proyecto sigue una arquitectura limpia en capas (**Handler** → **Service** → **Repository**) y utiliza Go Modules para la gestión de dependencias.

## 🎯 Funcionalidades Clave

La API soporta las operaciones CRUD (Crear, Leer, Actualizar, Eliminar) en las siguientes entidades:

### 1. Dashboard (Tableros)

* **Contenedor Principal:** Un usuario puede crear y gestionar múltiples tableros.
* **Inicialización de Flujo:** Al crear un nuevo `Dashboard`, automáticamente se inicializa un `Estado` por defecto llamado **"Pendiente"** para empezar el flujo de trabajo.

### 2. Estado (Columnas/Flujos)

* **Personalización:** Dentro de cada `Dashboard`, el usuario puede definir los estados (columnas) por los que pasará una `Tarea`.
* **Ejemplos de Estados:** "Pendiente" (por defecto), "En Progreso", "Revisión", "Completada".

### 3. Tarea (Elementos de Trabajo)

* **Creación:** Las tareas son creadas dentro de un `Dashboard` específico.
* **Asignación de Estado:** Cada tarea es asignada a uno de los `Estado` definidos en el `Dashboard`.
