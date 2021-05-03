var currentExhibitionID;

jQuery(function(){
    initialExhibition();
});


async function initialExhibition() {
    try {
        exhibitionList = await getExhibitionList();

        if(exhibitionList.length > 0) {
            updateExhibitionTitleUI(exhibitionList[0].title);
            updateExhibitionDropDownListUI(exhibitionList);

            updateExhibitionUI(exhibitionList[0].id);
        } else {
            updateExhibitionTitleUI("none");
        }
    } catch(error){
        console.log(error);
    }
}

function updateExhibitionTitleUI(title) {
    let titleUI = $("#exhibition-title").text(title);
}

function updateExhibitionDropDownListUI(exhibitionList) {
    let dropListUI = $("#exhibition-drop-down-list");

    dropListUI.empty();
    for(let exhibitionInfo of exhibitionList) {
        let dropListItemUI = $(`<li class="dropdown-item">${exhibitionInfo.title}</li>`);
        dropListItemUI.on('click', function(){
            updateExhibitionTitleUI(exhibitionInfo.title);
            updateExhibitionUI(exhibitionInfo.id);
        });

        dropListUI.append(dropListItemUI);
    }
}

async function updateExhibitionUI(exhibitionID) {
    try {
        exhibitionData = await getExhibitionData(exhibitionID);
        updateExhibitionPaintSlide(exhibitionData["paints"]);
    } catch(error) {
        console.log(error);
    }
}

function updateExhibitionPaintSlide(paintList) {
    let paintSlideUI = $("#exhibition-paint-slide");

    paintSlideUI.empty();
    
    for(let paintData of paintList) {
        paintSlideUI.append(`<div class="carousel-item">
            <img src="image/${paintData.url}" class="d-block mx-auto" alt="...">
        </div>`);
    }

    if (paintSlideUI.children().length > 0) {
        let firstPaint = paintSlideUI.children().eq(0);
        firstPaint.addClass("active");
    }
}

function getExhibitionList() {
    return new Promise(function(resolve, reject){
        $.ajax({
            url: "api/exhibition/list",
            dataType: "json"
        }).done(function(result) {
            exhibitionList = result["data"];
    
            resolve(exhibitionList);
        }).fail(function(){
            reject("call exhibition list api error");
        });
    });
}

function getExhibitionData(exhibitionID) {
    return new Promise(function(resolve, reject){
        $.ajax({
            url: `api/exhibition?id=${exhibitionID}`,
            dataType: "json"
        }).done(function(result) {
            exhibitionData = result["data"];
            resolve(exhibitionData);
        }).fail(function(){
            reject("call exhibition api error");
        });
    });
}
