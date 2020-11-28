function NewApartmentView() {
  function showLocations(element, locations) {
    const markup = `
            ${locations.map((location) => apartmentItem(location)).join("")}
        `;
    element.innerHTML = markup;
  }

  function showApartmentLocation(element, location) {
    element.innerHTML = apartmentLocationItem(location);
  }

  function showApartmentHeaderItem(element, data) {
    element.innerHTML = apartmentHeaderItem(data);
  }

  return {
    showLocations,
    showApartmentLocation,
    showApartmentHeaderItem
  };
}

const apartmentItem = (item) => `
 <div class="vehiculs-box">
 <div class="vehiculs-thumb">
     <img src="${getImageURL(item)}" alt="" />
     <span class="spn-status"> ${item.status} </span>
     <span class="spn-save"> <i class="ti ti-heart"></i> </span>
     <a class="proeprty-sh-more" href="location.html?id=${item.uuid}"><i
             class="fa fa-angle-double-right"> </i><i
             class="fa fa-angle-double-right"> </i></a>
     <p class="car-info-smal">
     ${item.description}
     </p>
 </div>
 <h3><a href="location.html?id=${item.uuid}" title="${item.name}">${item.name}</a></h3>
 <span class="price">${item.price.toLocaleString('pt-BR', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2
  })}</span>
</div>
`;

const getImageURL = (item) => {
  if (item.images !== null && item.images.length > 0) {
    return item.images[0].url;
  }
  return "img/demo/vehicul1.jpg"
}

const apartmentHeaderItem = (item) => `
<div class="inner-head overlap">
<div style="background: url(${getImageURL(item)}) repeat scroll 50% 422.28px transparent;"class="container">
    <div class="inner-content">
        <span><i class="ti ti-home"></i></span>
        <h2><span style="color: black; "> ${item.name}</span></h2>
        <ul>
                        <li><a  style="color: black; " href="#" title="">HOME</a></li>
                        <li><a  style="color: black; " href="#" title="">The Helux villa</a></li>
                    </ul>
    </div>
</div>
</div>
`
const apartmentLocationItem = (item) => `

<div class="blog-post vehicul-post">

  <div class="vehicul-gallery">
    <div class="light-slide-item">
        <div class="favorite-and-print">
            <ul id="image-gallery" class="gallery list-unstyled cS-hidden">
                ${item.images.map(image =>
                  `
                <li data-thumb="${image.url}">
                  <img src="${image.url}" alt="${image.file_name}" />
               </li> 

               `).join('')}
            </ul>
        </div>
    </div>
</div>

<h1>Hotel-fazenda hospedado por Carlos Hugo</h1>
<div class="hospedes">2 hóspedes · 1 cama · 1 banheiro</div>
<br/>

<div class="row">
    <div class="col-md-12">
        <p>
        ${item.description}
        </p>
    </div>
</div>
<div class="vehicul-feature">
                                            <div class="heading3">
                                                <h2>Comodidades </h2> 
                                            </div>
                                            <div class="vehicul-feature-content clearfix">
                                                <div class="has">
                                                    <i class="fa fa-check-circle"></i> attic</div>
                                                <div class="has">
                                                    <i class="fa fa-check-circle"></i> gas heat</div>
                                                <div class="no-has">
                                                    <i class="fa fa-times-circle"></i> balcony</div>
                                                <div class="no-has">
                                                    <i class="fa fa-times-circle"></i> wine cellar</div>
                                                <div class="has">
                                                    <i class="fa fa-check-circle"></i> basketball court</div>
                                                <div class="has">
                                                    <i class="fa fa-check-circle"></i> trash compactors</div>
                                                <div class="has">
                                                    <i class="fa fa-check-circle"></i> gym</div>
                                                <div class="no-has">
                                                    <i class="fa fa-times-circle"></i> fireplace</div>
                                                <div class="has">
                                                    <i class="fa fa-check-circle"></i> pool</div>
                                                <div class="no-has">
                                                    <i class="fa fa-times-circle"></i> lake view</div>
                                                <div class="has">
                                                    <i class="fa fa-check-circle"></i> solar Heat</div>
                                                <div class="has">
                                                    <i class="fa fa-check-circle"></i> storm Windows</div>
                                                <div class="has">
                                                    <i class="fa fa-check-circle"></i> separate Shower</div>
                                                <div class="no-has">
                                                    <i class="fa fa-times-circle"></i> wet bar</div>
                                                <div class="no-has">
                                                    <i class="fa fa-times-circle"></i> remodeled</div>
                                                <div class="no-has">
                                                    <i class="fa fa-times-circle"></i> skylights</div>
                                                <div class="has">
                                                    <i class="fa fa-check-circle"></i> stone Surfaces</div>
                                                <div class="no-has">
                                                    <i class="fa fa-times-circle"></i> open entertaining Kitchen</div>
                                                <div class="has">
                                                    <i class="fa fa-check-circle"></i> golf course</div>
                                                <div class="no-has">
                                                    <i class="fa fa-times-circle"></i> health club</div>
                                                <div class="no-has">
                                                    <i class="fa fa-times-circle"></i> backyard</div>
                                                <div class="has">
                                                    <i class="fa fa-check-circle"></i> pet allowed</div>
                                                <div class="has">
                                                    <i class="fa fa-check-circle"></i> office/den</div>
                                                <div class="has">
                                                    <i class="fa fa-check-circle"></i> laundry</div>
                                            </div>
                                        </div>
</div><!-- Blog Post -->
`