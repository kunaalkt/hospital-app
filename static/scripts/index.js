document.getElementById("loginForm").addEventListener("submit", function(e) {
  e.preventDefault();

  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;

  fetch("http://localhost:8080/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      username,
      password
    })
  })
  .then(response => response.json())
  .then(data => {
    const resDiv = document.getElementById("response");
    if (data.token) {
      localStorage.setItem("token", data.token);
      localStorage.setItem("role", data.role);

      if (data.role === "receptionist") {
        window.location.href = "/receptionist.html";
      } else if (data.role === "doctor") {
        window.location.href = "/doctor.html";
      } else {
        resDiv.innerText = "Unknown role";
      }
    } else {
      resDiv.style.color = "red";
      resDiv.innerText = "Login failed: " + (data.error || "Invalid credentials");
    }
  })
  .catch(error => {
    document.getElementById("response").innerText = "Error: " + error;
  });
});