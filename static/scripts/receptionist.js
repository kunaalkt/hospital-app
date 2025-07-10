const apiBase = 'http://localhost:8080/receptionist';
const token = localStorage.getItem("token");
if (!token) {
  window.location.href = '/';
}

async function fetchPatients() {
  const res = await fetch(apiBase + '/patients', {
    headers: { 'Authorization': 'Bearer ' + token }
  });
  const patients = await res.json();
  const ul = document.getElementById('patients-list');
  ul.innerHTML = '';
  patients.forEach(p => {
    const li = document.createElement('li');
    li.innerHTML = `
      <div class="mt-3 p-4 border border-3">
        <p>ID: ${p.ID}</p>
        <p>Name: ${p.Name}</p>
        <p>Age: ${p.Age}</p>
        <p>Gender: ${p.Gender}</p>
        <p>Diagnosis: ${p.Diagnosis}</p>
        <p>Prescription: ${p.Prescription}</p>
        <button id="update-button" class="btn btn-primary" onclick="updatePatientPrompt(${p.ID})">Update</button>
        <button id="delete-button" class="btn btn-danger ms-1" onclick="deletePatient(${p.ID})">Delete</button>
      </div>
    `;
    ul.appendChild(li);
  });
}

async function createPatient() {
  const name = document.getElementById('name').value;
  const age = +document.getElementById('age').value;
  const gender = document.getElementById('gender').value;
  const diagnosis = document.getElementById('diagnosis').value;
  const prescription = "N/A";
  const res = await fetch(apiBase + '/patients', {
    method: 'POST',
    headers: { 
      'Authorization': 'Bearer ' + token,
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ name, age, gender, diagnosis, prescription })
  });
  if(res.ok) {
    alert('Patient created');
    fetchPatients();
    } else {
    alert('Failed to create patient');
  }

  resetForm();
}

async function updatePatientPrompt(id) {
  const newName = prompt('New name:');
  const newAge = prompt('New age:');
  const newGender = prompt('New gender:');
  const newDiagnosis = prompt('New diagnosis:');
  const newPrescription = "N/A";
  if(newName && newAge && newGender && newDiagnosis) {
    await updatePatient(id, newName, +newAge, newGender, newDiagnosis, newPrescription);
  } else {
    alert('Invalid patient details!');
  }
}

async function updatePatient(id, name, age, gender, diagnosis, prescription) {
  const res = await fetch(apiBase + '/patients/' + id, {
    method: 'PUT',
    headers: {
      'Authorization': 'Bearer ' + token,
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ name, age, gender, diagnosis, prescription })
  });
  if(res.ok) {
    alert('Patient updated');
    fetchPatients();
  } else {
    alert('Failed to update patient');
  }
}

async function deletePatient(id) {
  const res = await fetch(apiBase + '/patients/' + id, {
    method: 'DELETE',
    headers: { 'Authorization': 'Bearer ' + token }
  });
  if(res.ok) {
    alert('Patient deleted');
    fetchPatients();
  } else {
    alert('Failed to delete patient');
  }
}

function logout() {
  const conf = confirm("Are you sure you want to logout?");
  if (!conf) return;

  localStorage.removeItem('token');
  localStorage.removeItem('role');

  window.location.href = '/';
}

function resetForm() {
  document.getElementById('name').value = "";
  document.getElementById('age').value = "";
  document.getElementById('gender').value = "";
  document.getElementById('diagnosis').value = "";
}

fetchPatients();