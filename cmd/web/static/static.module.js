export let checkAvailabilityStaticForm = `
  <form id="check-availability-form" action="" method="post" novalidate class="needs-validation">
      <div class="form-row">
          <div class="col">
              <div class="form-row" id="reservation-dates-modal">
                  <div class="col">
                      <input disabled required class="form-control" type="text" name="start_date" id="start" placeholder="Arrival">
                  </div>
                  <div class="col">
                      <input disabled required class="form-control" type="text" name="end_date" id="end" placeholder="Departure">
                  </div>
              </div>
          </div>
      </div>
  </form>
`;
