Codes = {
    black = 0,
    brown = 1,
    red = 2,
    orange = 3,
    yellow = 4,
    green = 5,
    blue = 6,
    violet = 7,
    grey = 8,
    white = 9
}

return {
    decode = function(c1, c2, c3)
        c1 = Codes[c1]
        c2 = Codes[c2]
        c3 = Codes[c3]

        local value = (10 * c1 + c2) * 10 ^ c3
        if (value / 1000 > 1) then
            return value / 1000, "kiloohms"
        else
            return value, "ohms"
        end
    end
}
