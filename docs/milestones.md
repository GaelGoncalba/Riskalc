# Riskalc - Milestones del Proyecto

### [M0] Milestone 0: Definición del problema
- **Descripción:** Definir los conceptos clave del dominio del problema sobre la unificación de criterios del cálculo de riesgos, así como las diferentes entidades que forman parte del mismo y sus relaciones. 
- **Entregable:** Modelo inicial del problema recogido, con los conceptos clave y sus relaciones, definiendo las entidades y agregados necesarios para desarrollar el PMV. Se entregará en forma de código fuente.  
- **Viabilidad:** Se considerará validado cuando el equipo de desarrollo y revisión comprenda y defina los elementos del dominio del problema, recogido en las historias de usuario.

---

### [M1] Milestone 1: Unificación de criterios de cálculo de riesgo
- **Descripción:** Implementar un mecanismo capaz de comparar los parámetros de un nuevo préstamo con los de las operaciones recogidas en el historial, para asignarle un grupo de "operaciones considerablemente similares".
- **Entregable:** Código encargado de realizar el algoritmo de clustering.
- **Viabilidad:** Se realizará un análisis mediante test para asegurar que el nivel de coincidencia de las operaciones pasadas con la nueva supera el umbral establecido.

---

### [M2] Milestone 2: Automatización del cálculo del riesgo
- **Descripción:** Implementar la automatización del cálculo del riesgo a partir de las operaciones introducidas en el cluster asignado, estudiando cuántas de ellas resultaron pagadas.
- **Entregable:** Versión funcional del código con la lógica de cálculo de riesgos integrada.
- **Viabilidad:** Se considerará validado cuando se pasen los test propuestos.

---

## Milestones adicionales:

### [M3] Milestone 3: Implementar el filtrado de operaciones en el historial 
