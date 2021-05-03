var currentExhibitionIndex = null;
var storedExhibitionList = [];
var lastClickTableRow = null;
var lastClickExhibitionIndex = null;
var bootstrapToast = null;

jQuery(function(){
    initial();
});

async function initial() {
    try {
        let toastUI = $('#status-toast').get()[0];
        let option = {
            animation: true,
            autohide: true,
            delay: 2000
        };
        bootstrapToast = new bootstrap.Toast(toastUI, option);

        let saveExhibitionButton = $('#save-exhibtion-setting-btn');
        saveExhibitionButton.on('click', createExhibition);

        storedExhibitionList = await getExhibition();

        if (storedExhibitionList.length > 0 ) {
            updateExhibitionTableUI(storedExhibitionList);
        }
        
        $("#create-exhibition-btn").on("click", function(){
            var createExhibitionModel = new bootstrap.Modal(document.getElementById('create-exhibition-modal'), {
                keyboard: false
            });

            createExhibitionModel.show();
        });

        $("#add-paint-btn").on("click", function(){
            var addPaintModel = new bootstrap.Modal(document.getElementById('add-paint-modal'), {
                keyboard: false
            });

            
            if (lastClickTableRow == null ) {
                return;
            }

            updateAddPaintModelUI();
            addPaintModel.show();
        });

        $("#add-paint-modal button.btn-danger").eq(0).on("click", addPaintToExhibition);
    } catch( error ) {
        console.log(error);
    }
}

function getExhibition() {
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

function createExhibition() {
    let exhibitionTitleInput = $("#exhibition-title-input");
    let exhibitionDescriptionInput = $("#exhibition-description-input");

    let requestPayload = {
        title: exhibitionTitleInput.val(),
        description: exhibitionDescriptionInput.val()
    };

    $.ajax({
        url: "api/exhibition",
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify(requestPayload),
        dataType: "json"
    }).done(function(result) {
        storedExhibitionList.push(result['data']);

        updateExhibitionTableUI(storedExhibitionList);
        updatePaintTableUI([]);
    }).fail(function(){
        showToast("Create exhibition error");
    });
}

function updateExhibitionTableUI(exhibitionList) {
    let exhibitionTableUI = $("#exhibition-table");

    lastClickTableRow = null;
    exhibitionTableUI.empty();

    for(let index in exhibitionList) {
        let exhibitionTableRow = $(`<tr>
            <td>${exhibitionList[index].id}</td>
            <td>${exhibitionList[index].title}</td>
            <td>${exhibitionList[index].description}</td>
        </tr>`);
        
        exhibitionTableRow.on('click', function(){
            if (lastClickTableRow != null) {
                lastClickTableRow.removeClass("table-success");
            }

            exhibitionTableRow.addClass("table-success");
            updatePaintTableUI(storedExhibitionList[index].paints);
            lastClickExhibitionIndex = index;
            lastClickTableRow = exhibitionTableRow;
        });

        exhibitionTableUI.append(
            exhibitionTableRow
        );
    }
}

function updatePaintTableUI(paintList){
    let paintTableUI = $("#paint-table");

    paintTableUI.empty();
    if (paintList) {
        for(paint of paintList) {
            paintTableUI.append(
                `<tr>
                    <td>${paint.id}</td>
                    <td>${paint.name}</td>
                </tr>`
            );
        }
    }
}

function getPaintList() {
    return new Promise(function(resolve, reject){
        $.ajax({
            url: "api/paint/list",
            dataType: "json"
        }).done(function(result) {
            paintList = result["data"];
    
            resolve(paintList);
        }).fail(function(){
            reject("call paint list api error");
        });
    });
}

async function updateAddPaintModelUI() {
    try {
        let availablePaintTableUI = $("#available-paint-table");
        availablePaintTableUI.empty();

        let paintList = await getPaintList();
        for( let paint of paintList ) {
            let paintRecord = `
                <tr>
                    <td>
                        <input type="checkbox">
                    </td>
                    <td>${paint.id}</td>
                    <td>${paint.name}</td>
                </tr>
            `;

            availablePaintTableUI.append(paintRecord);
        }
    } catch(error) {
        console.log(error);
        $('#status-toast div.toast-body').text("Error on adding paint to exhibition");
        bootstrapToast.show();
    }
}

async function addPaintToExhibition() {
    try {
        let paintRecordList = $("#available-paint-table tr");
        let checkedRecord = [];

        for(let paintRecord of paintRecordList) {
            let checkbox = $(paintRecord).find(`input[type="checkbox"]`);
            if(checkbox.prop("checked")) {
                let paintIdData = $(paintRecord).find("td").eq(1).text();
                let paintId = parseInt(paintIdData);

                checkedRecord.push(paintId);
            }
        }
        
        exhibitionId = parseInt(lastClickTableRow.children("td").eq(0).text());
        callAddPaintToExhibitionApi(exhibitionId, checkedRecord);
    } catch(error) {
        console.log(error);
        $('#status-toast div.toast-body').text("Error on adding paint to exhibition");
        bootstrapToast.show();
    }
}

function callAddPaintToExhibitionApi(exhibitionId, paintIdList) {
    let requestPayload = {
        exhibition_id: exhibitionId,
        paint_id_list: paintIdList
    };

    console.log(JSON.stringify(requestPayload));
    return new Promise(function(resolve, reject){
        $.ajax({
            url: "api/exhibition/paint",
            method: "POST",
            contentType: "application/json",
            data: JSON.stringify(requestPayload),
            dataType: "json"
        }).done(function(result) {
            if( result.status == 'ok') {
                showToast("Add paint to exhibition success");
                updatePaintTableUIByExhibitionID(exhibitionId);
            } else {
                showToast("Add paint to exhibition failed");
            }    
        }).fail(function(){
            showToast("Add paint to exhibition failed");
        });
    });
}

async function updatePaintTableUIByExhibitionID(exhibtionId) {
    try {
        let paintList = await getPaintListByExhibitionId(exhibtionId);
        storedExhibitionList[lastClickExhibitionIndex].paints = paintList;
        updatePaintTableUI(paintList);
    } catch(error) {
        console.log(error);
    }
}

function getPaintListByExhibitionId(exhibtionId) {
    return new Promise(function(resolve, reject){
        $.ajax({
            url: `api/exhibition/paint/list?exhibition_id=${exhibtionId}`,
            dataType: "json"
        }).done(function(result) {
            paintList = result["data"];
    
            resolve(paintList);
        }).fail(function(){
            reject("call exhibition paint list api error");
        });
    });
}

function showToast(message) {
    $('#status-toast div.toast-body').text(message);
    bootstrapToast.show();
}




