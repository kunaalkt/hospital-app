<div id="readme-top" align="center">
    <h1>Hospital App</h1>
    <h4>A simple web-based application using Golang to register and manage patients of a hospital</h4>
    </br>
</div>

<div id="table-of-contents">
    <details>
    <summary>Content List</summary>
    <ul>
        <li>
        <a href="#about-the-project">About The Project</a>
        <ul>
            <li><a href="#built-with">Built With</a></li>
        </ul>
        </li>
        <li>
        <a href="#getting-started">Getting Started</a>
        <ul>
            <li><a href="#requirements">Requirements</a></li>
            <li><a href="#installation">Installation</a></li>
        </ul>
        </li>
        <li><a href="#endpoints">Endpoints</a></li>
    </ul>
    </details>
</div>

<div id="about-the-project">
    <h2><u>About the Project</u></h2>
    <p>This project is a web-based application developed using the Go programming language. The system provides two user roles: Receptionist and Doctor, each with distinct functionalities, all accessible through a unified login page.</p>
    <h3>Key Features:</h3>
    <ol>
        <li><b>Authentication:</b></li>
        <ul>
            <li>A single login endpoint (/login) supports both Receptionists and Doctors</li>
            <li>Role-based access control is implemented to differentiate functionalities for each user type</li>
        </ul>
        <li><b>Receptionist Portal:</b></li>
        <ul>
            <li>Register new patients</li>
            <li>View patient list</li>
            <li>Edit patient information</li>
            <li>Delete patient records</li>
            <li>Full CRUD operations on patient data</li>
        </ul>
        <li><b>Doctor Portal:</b></li>
        <ul>
            <li>View all registered patients</li>
            <li>Access detailed medical/personal data</li>
            <li>Prescibe patient-related on the go</li>
        </ul>
    </ol>
    <div id="built-with">
        <h3>Built with:</h3>
        <ul>
            <li>Frontend Technologies: HTML, CSS, Bootstrap, JavaScript</li>
            <li>Backend Technologies: Go, Gin, GORM, Postman</li>
            <li>Database: PostgreSQL</li>
        </ul>
    </div>
</div>

<div id="getting-started">
    <h2><u>Getting Started</u></h2>
    <div id="requirements">
        <h3>Requirements:</h3>
        <ul>
            <li>Tools: Visual Studio Code, Go Compiler, Postman</li>
            <li>Internet Connection</li>
        </ul>
    <div id="installation">
        <h3>Installation: </h3>
        <ol>
            <li>Download/Clone the project repository</li>
            <li>Setup your environment variables (.env) in the root folder for DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT</li>
            <li>Install all the requirements from go module using <code>go mod tidy</code></li>
            <li>Create a database (hospital_db) and connect to it. Create a table (users) and create your doctor and receptionist accounts.</li>
            <li>Run the go app in terminal using <code>go run cmd/main.go</code></li>
        </ol>
    </div>
    </div>
</div>

<div id="endpoints">
    <h2><u>Endpoints</u></h2>
    <ul>
        <li>
            <code>POST /login</code> – login, get JWT
        </li>
        <li>
            <code>POST /patients</code> – receptionist only
        </li>
        <li>
            <code>DELETE /patients/:id</code> – receptionist only
        </li>
        <li>
            <code>GET /patients</code> – receptionist and doctor
        </li>
        <li> <code>GET /patients/:id</code> – receptionist and doctor</li>
        <li><code>PUT /patients/:id</code> – receptionist and doctor</li>
    </ul>
</div>