<template>
    <div class="chat-input-wrapper">
        <img  src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADgAAAA6CAYAAADlTpoVAAAAAXNSR0IArs4c6QAACa5JREFUaEPNmgtwlNUVx39nd5OQB+EZXgqCEQGjFauILyhWbVHAlk6JQuUNVQio42Ns67STadWx4GOEJGgBIziChtpWaRFLrTwr+BYbAyISeUUJIOTNZvc7nfvtLgLZTfbb/Wx7ZnYDs/f8z/nf57nnXMFlGVem3sxqvmspw1AGioeBKL2B9uGPsVhrf4R9KDsVKjzCpvoc3l+VL0E3XRI3wH7+jKY0BhjpsRgPXI/QAfAABj/0Md9qf5v/KWr/y3xHPhbKcYQ3LGFFuo+1f7hdmpP1LymC457Q9Kw0CizlToQcFB+C9yQRp96FiAcRAijVKAsamiledY80OoWKtE+IoJmG2UeYEAjyK6AvkBYeqUT9iKZnRvYEUCnwSH03ViQyfR0TnFai11sWhQoXAxnhqegmsTOxLIUGYHuqj8Ild8g6J8biJ6gqUxfxC8vi/vBm4XNiKNm2CgGBWo+H+aWzeBSR0CpuQ+IiOHG+ZnozeEqVcUB2W6Df8u81IqwKNnDX8/dLfVu22iQ4tVh7q7BYLYaFp2RbmC1+z2wHXbIgq13op7omOFIH9U2OoSIKDeJhkygzSwtkX2sorRKcsVjPDvj5E8p3NLSRxC1pPhjaHy7sA2d1Bp8HPGFrZlEFLDj4Nfx7L2zbBU0ODwSBEwofp6QxdslM2R/LsZgEJy7XTG8tZar8AHC03gb3hZsGQ9dsaJcKKd7o5puD0OSHw7Xw2ofwwZ64+y/SMCDC34/4GLf6djEbUQuJTtBsKE+zwLKYgRKeWPEZv3EwjMiDjpnfjFhbmmZEj9fDhgpY837o5I9bhCYPLCmdbc7ilhtPVILTSrQgaPEQ0DFuQ8CwgTD2cmif7kTrm7a1TfDKO7DhE8f6x3zCr5cWSNGZmi0I3rFIr24KUqbQy4mZ3l1hzkjonOlEq2XbY/VQ/DpUVjvDETjYzkv+07Nky6mapxEsVPVUlvCqKqOcwcP078MV/Z1qRW//3ufwtKPjPIQjwpq+sxlTKGJFkE8jOHORjvYHeRFwNA49O8Ivx0J6qjsEzY467xXYd8QxXr3Hy62ls+SvLQjOXaBpdV7W2dcch3LdRXDrVQ6V2mj+8jZY+6FzTHPt+jLIDa/dKSaODV9fgOklOiGgLHW6axoQs/YuPse5M61pVByAJ06OgyPsJlGmPzdXVpwkWPim+irL2aJwuSOocONHxkOOywHc8Qa47/lEvLFH7Z36blxpbh/2GpyxWK9obmIDguNVZAAWToO0lMSciaUVtKBgKZi/jkXxe1MY8ewd8pZNcFqJPh60uMcxkAlxvFAy3d7BXJc7S6HRnxisR3iytEDukcJC9VR25ROFAYlAmfjSjGCqo2CubUtm5OY+CyacS0RE2Nm3mgtkerGeF1DKwfn0jBh+dAJ0MSklF6WmEe5dnhSg3yfkydQSvc2ySHA5hxx44EdwXo+knGmhvPcw/O7l5DBVmShTi3W+pdyXDNTkEXBNQhM8ttVtn8GSN5Lxyj4EH5MpxWpCszHJQJkQzYRqbsqyDbB5R3KIIqyWKUW6XeGiZKA6ZsBvb3EvVDsRgN+8BEfrkvHKPg8/lskL9SBCz+Sg4O5RkHd2sigh/Z1V8PirDu+F0UwrVTK5SE0aPStZ14bkwozr4r/kxrJnLr+l62Hrp8l6ZOvXGYImKHUcwZxpPiMV7roJzu2enGN7DsFTa6DeDpWTFr9rI2hcMaNoNhuvqUokIOZwf3Y9vL0rAeXoKnWurUGDb+LR6dfCJf0Sc/DDSlj6ZigR5YqE1+B2ktxFT3Wmbw5M/B706eLMxf1HYPlGMFPURfnYTNFXgJtdBOWyXBh/NWTHmXwyieCVW+Dtz9z0wsZ6VSYV6TzBrje4JiYAnzICrjw/Psj3P4dn3gArkatRKyYU5sukhTpBhBficyX+VobctGvja79yM/zThPsuiyo/k2klem7QosKNo+JU/4YNgknD4/PYjbAsiiW/18Mg+z64p6t9XRoYnzttt8pMAxOAX2JKo3HIpgp4YXOCt/fY+Dv6HSbPvodPLtLHgHvj8KXVJoZY9w4woBeMvjT+S7BJ9q7aCgeOQnUN+APJemLrP75sjtxnE5xapEMtZWNCORmBTpkhYufkwIW9Q3dDp4e9iVx27IeKg6GqkyFqEk/qqFAR7hjF7xGGl86RbTZBU3PPqGYzyhVO+s7U/EzK3ozYRX1CJJMVQ2jvEfhkP3xWFSJraomOiApbG3K45mRWzTg1aaHeIsJz0HY1yRQye3UOrbGr+kOGo/pT/F1Q2wgmjf/RF6EstxnROKRJlSnL58pLpu3JXNi4Mk3NOMTrwIhYICaD1qsTXHA2DB/kfi40ll2Tn9lYAR9Vwv6jEGg9EbW+oRs/XJUvdsB3WrJv0kK90dS/o9UmMtKgfw8YOdj9/Esco2I3MRWnNR/ArqpQGTyK1Jt3BMvnymuR306vLpkjI4e/cEYKw0zJoeeFdsZInT1ep9xuZ/KkhuS/doIZ2dNEWN2vmh8XFsaoLtlrsUivBMoE7Pu5qRhdeyHcfKnzndFtchE8E9KZkve67d/cGxVMnT5/+Rx561S7UfPRk4t0FvCwR+h09QC4bXjyN3W3yZrT48UtsL4cLOVr4MFlc2TRmXZi1uinlPBklyxmP/gTUv7X0zJW5zT44aGXaa6uYdGyAu6Ou0ZvAM0ri9Hn8o8xQxjq85y+Gbk9GoniBSx0zbts27qb6x+bFP1RUKslkz++o5cM6ctLXduTm+K1n0f+30hzEOtQDbvf+4JbfjpEPojlWJs1oeJ1OuiGC1h1VhcGpHidvZf5tnqjOUjgwFF2rCsnv+AGMTehmNImQaNZskk7DevD2tzuXOzzOnvx5DbJQJATu7/io017GTl7mJjNpVWJi6BBKFP15u7hz4POYrhXyJbQG97/mqiiAaVmxwE27u7H2HyJ7+mzYyf/9q7edVku92dn0NUjpH0bhc9Te80E2ZZyoqaBw+/uZt6oy2SBk151TNCAlx/SLH8ti3J7MjLFSwePkOI20TCx5uYgx3dXsTa1PbPyuonjakVCBCM9WH5Ie9QcY3H/7lyTlY7JoRmikUeFTjrabmu/TldM6qm5rpHGXV+xObsjM/O6yZeOwcIKSRGMGH1mtWb0680DPTswvnc3eqankqKKT0Kv7c1Tc3vFRozZd1jzFXqEbgZLRQg0+mned4iqquOs3LOP398+JvoLQidkXSF4qsGyck1t72d8ZipjMtM5v2MGPbLakZHixZueGuLY6EebgwTrmmg41sCX9Y18Wu9ndW0qK/PzQtcct8R1gm455hbOfwCaMUFitnGvRwAAAABJRU5ErkJggg==" alt="">
        <div class="chat-me">{{content}}</div>
    </div>
    
</template>

<script>
export default {
    name: "WindowUser",
    props: {
        content: { type: String, required: true },
    },
    data(){
        return{
        }
    },
    mounted() {
    },
}
</script>

<style scoped lang="less">
.chat-input-wrapper{
    margin-left: 60px;
    margin-bottom: 24px;
    display: flex;
    flex-direction: row-reverse;
    white-space: pre-line;
    img{
        width:32px;
        height:32px;
        margin-left: 10px;
    }
    .chat-me{
        display: inline-block;
        background-color:#ebebfb;
        border-radius: 6px;
        padding: 8px;
        line-height: 20px;
        white-space: pre-wrap;
    }
}
@media (max-width: 768px){
    .chat-input-wrapper{
        margin-left: 32px;
    }
    
}
</style>