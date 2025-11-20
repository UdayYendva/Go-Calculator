# Go-Calculator
This Calculator Application is build using GO and implemented CRUD operations to take user inputs like POST GET DELETE and UPDATE
💻 README: Go Calculator with POST JSON InputThis document provides an overview of the Go Calculator API service designed to perform basic arithmetic operations using a POST request that accepts input in JSON format.
🚀 Getting StartedPrerequisitesTo run this service, you need:Go (version 1.18 or higher) installed on your system.Installation and RunningClone the repository (Assuming your Go calculator code is in a repository).Navigate to the project directory.Run the application:Bashgo run main.go
(Note: Replace main.go with your actual entry file if different.)The server will typically start on http://localhost:8080, or the port you have configured.
📝 API EndpointThe calculator service exposes a single endpoint for all operations.Endpoint DetailsMethodPathDescriptionPOST/calculatePerforms the arithmetic calculation.⚙️ Input Format (JSON)The service only accepts input via the request body in JSON format using the POST method.JSON StructureThe JSON body must contain the following three fields:FieldTypeDescriptionExampleoperand1float64 (or number)The first number for the calculation.10.5operand2float64 (or number)The second number for the calculation.5operatorstringThe arithmetic operator to use."+"Supported OperatorsOperatorOperation+Addition-Subtraction*Multiplication/DivisionExample Request BodyJSON{
    "operand1": 25.5,
    "operand2": 5.0,
    "operator": "/"
}
✅ Output Format (JSON)The service responds with a JSON object containing the result or an error message.Success ResponseFieldTypeDescriptionresultfloat64 (or number)The calculated result.Example Success Response (for the request above):JSON{
    "result": 5.1
}
Error ResponseFieldTypeDescriptionerrorstringA message detailing the error.Example Error Response (e.g., for division by zero):JSON{
    "error": "Division by zero is not allowed."
}
Example Error Response (e.g., for invalid operator):JSON{
    "error": "Invalid operator: ^. Supported operators are +, -, *, /."
}
