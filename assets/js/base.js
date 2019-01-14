$(document).ready(function () {
    initApp();

    $(".clouds").on("click", function () {
        openDiaryModal();
        $('[data-toggle="popover"]').each(function () {
            $(this).popover("hide");
        });
    });

    $("#add_diary").on("hidden.bs.modal", function () {
        doOpenWelcomeStar();
        randomDisplayQuote(false);
    });

    // randomCloud();

    randomDisplayQuote(false);
});

var t = null;

function sendContent() {
    var value = $.trim($("#content-diary").val());
    $.ajax({
        type: "POST",
        url: '/api/diaries',
        data: {
            'Content': value,
        }, success: function (data) {
            if (!data.status) {
                swal("Oops!", data.err.Message, "warning");
            } else {
                swal("Great!", "Your stories are safe with the stars!!!", "success").then((value) => {
                    if (value) {
                        location.reload();
                    }
                });
            }
        }
    });
}

function randomIntFromInterval(min, max) {
    return Math.floor(Math.random() * (max - min + 1) + min);
}

function doSomeStuff(data) {
    $('[data-toggle="popover"]').popover({
        template: '<div class="popover bg-primary"><div class="arrow"></div><h3 class="popover-title"></h3><div class="popover-content"></div></div>'
    });

    setTimeout(function () {
        $("#slogan").fadeOut(3000);
        $("#container-star").removeClass("hidden");

        setTimeout(function () {
            doOpenWelcomeStar();
        }, 3000);
    }, 2000);
}

function renderStar(data) {
    $.each(data, function (i, item) {
        $("#container-star").append(createHTML(item));
    });
}

function createHTML(item) {
    var top = randomIntFromInterval(5, 85) + "%",
        left = randomIntFromInterval(5, 85) + "%",
        size = randomIntFromInterval(3, 6) + "px",
        positions = ["top", "right", "bottom", "left", "auto"],
        pos = positions[Math.floor(Math.random() * positions.length)];
    // triggers = ["click", "hover", "focus"],
    // trigger = triggers[Math.floor(Math.random()*triggers.length)];


    html = '<div style="top: ' + top + '; left: ' + left + '; width: ' + size + '; height: ' + size + '" class="star" id="star_' + item.ID + '" title="" data-toggle="popover" data-html="true" data-placement="' + pos + '" data-content="' + item.Content + '" data-trigger="click">';
    html += '</div>';
    return html;
}

function initApp() {
    $.ajax({
        type: "GET",
        url: '/api/diaries',
        success: function (data) {
            renderStar(data);
            doSomeStuff(data);
        }
    });
}

function doOpenWelcomeStar() {
    // open 4 stars on start
    $.each(new Array(4), function (n) {
        n = n + 1;
        $("#star_" + n).click();
    });
}

function openDiaryModal() {
    $("#add_diary").modal('show');
    doOpenWelcomeStar();
    $("#content-diary").val("");
    randomDisplayQuote(true);
}

function randomCloud() {
    var clouds = ["clouds4",],
        cloud = clouds[Math.floor(Math.random() * clouds.length)];
    $(".clouds").addClass(cloud);
}

function randomDisplayQuote(suspend) {
    if(suspend){
        clearInterval(t);
        return;
    }
    t = setInterval(function () {
        var len = $(".star").length;
        value = randomIntFromInterval(5, len);
        $("#star_" + value).click();
        setTimeout(function () {
            $("#star_" + value).click();
        }, 5000);
    }, 8000);
}