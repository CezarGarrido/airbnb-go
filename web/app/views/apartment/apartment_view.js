function NewApartmentView() {
  function showLocations(element, locations) {
    const markup = `
            ${locations.map((location) => apartmentItem(location)).join("")}
        `;
    element.innerHTML = markup;
  }

  return {
    showLocations: showLocations,
  };
}

const apartmentItem = (item) => `
 <div class="vehiculs-box">
 <div class="vehiculs-thumb">
     <img src="img/demo/vehicul1.jpg" alt="" />
     <span class="spn-status"> ${item.status} </span>
     <span class="spn-save"> <i class="ti ti-heart"></i> </span>
     <a class="proeprty-sh-more" href="vehicul.html"><i
             class="fa fa-angle-double-right"> </i><i
             class="fa fa-angle-double-right"> </i></a>
     <p class="car-info-smal">
     ${item.description}
     </p>
 </div>
 <h3><a href="vehicul.html" title="Mercedes-Benz">${item.name}</a></h3>
 <span class="price">${item.price.toLocaleString('pt-BR', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2
  })}</span>
</div>
`;
