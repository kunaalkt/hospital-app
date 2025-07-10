const apiBase = 'http://localhost:8080/doctor';
const token = localStorage.getItem("token");
const role = localStorage.getItem("role")
if (!token || !role) {
  window.location.href = '/';
}

async function fetchPatients() {
  const res = await fetch(apiBase + '/patients', {
    headers: {
      'Authorization': 'Bearer ' + token
    }
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
        <button id="button" onclick="updatePatientPrompt(${p.ID})" class="btn btn-primary">Update</button>
      </div>
    `;
    ul.appendChild(li);
  });
}

async function updatePatientPrompt(id) {
  const prescription = prompt('Prescription:');
  if (prescription) {
    await updatePatient(id, prescription);
  } else {
    alert('Invalid details!');
  }
}

async function updatePatient(id, pres) {
  const patientRes = await fetch(apiBase + '/patients/' + id, {
    headers: {
      'Authorization': 'Bearer ' + token
    }
  });

  if (!patientRes.ok) {
    alert('Failed to fetch patient data');
    return;
  }

  const patient = await patientRes.json();

  const updatedPatient = {
    ...patient,
    prescription: pres
  };

  const res = await fetch(apiBase + '/patients/' + id, {
    method: 'PUT',
    headers: {
      'Authorization': 'Bearer ' + token,
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(updatedPatient)
  });

  if (res.ok) {
    alert('Patient updated');
    fetchPatients();
  } else {
    alert('Failed to update patient');
  }
}

function logout() {
  const conf = confirm("Are you sure you want to logout?")
  if (!conf) return

  localStorage.removeItem('token')
  localStorage.removeItem('role')

  window.location.href = '/'
}

fetchPatients();