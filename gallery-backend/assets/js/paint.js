var storedPaintList = [];

jQuery(function(){
    initial();
});

async function initial() {
    try {
        storedPaintList = await getPaintList();
        updatePaintTableUI(storedPaintList);
    } catch( error ){
        console.log(error);
    }
};

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

function updatePaintTableUI(paintList) {
    let paintTableUI = $("#paint-table");

    for( let paint of paintList ) {
        let paintRecord = $(`<tr>
            <td>${paint.id}</td>
            <td>${paint.name}</td>
            <td><a href="image/${paint.url}">${paint.url}</a></td>
        </tr>`);

        paintRecord.on('click', function(){
            let name = $(this).find("td").eq(1).text();
            let url = $(this).find("td").eq(2).text();

            updatePaintPreviewUI(name,url);
        });

        paintTableUI.append(
            paintRecord
        );
    }
}

function updatePaintPreviewUI(name, url) {
    let paintPreviewUI = $(`#paint-preview`);

    let paintPreviewName = paintPreviewUI.find(`span`);
    let paintPreviewImage = paintPreviewUI.children(`img`);

    paintPreviewName.text(name);
    paintPreviewImage.prop('src', 'image/' + url);
}