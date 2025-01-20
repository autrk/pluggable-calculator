const std = @import("std");
const extism_pdk = @import("extism-pdk");
const Plugin = extism_pdk.Plugin;

const allocator = std.heap.wasm_allocator;

/// Greet function (bereits vorhanden)
export fn greet() i32 {
    const plugin = Plugin.init(allocator);
    const name = plugin.getInput() catch unreachable;
    defer allocator.free(name);

    const output = std.fmt.allocPrint(allocator, "Hello, {s}!", .{name}) catch unreachable;
    plugin.output(output);
    return 0;
}

/// Add function: Erwartet zwei float64-Werte im Input
export fn add() i32 {
    const plugin = Plugin.init(allocator);

    // Eingabe (Input) abrufen
    const input = plugin.getInput() catch unreachable;

    // Prüfen, ob genau 16 Byte (2 * float64) übergeben wurden
   // if (input.len != 16) {
   //     const error = "Invalid input length. Expected 16 bytes (2 float64).";
   //     plugin.output(error);
   //     return 1;
   // }

    // Zwei float64-Werte aus den 16 Byte lesen
    const a = std.mem.bytesToValue(f64, input[0..8]);
    const b = std.mem.bytesToValue(f64, input[8..16]);

    // Berechnung: a + b
    const sum = a + b;

    // Ergebnis als Byte-Array (float64) zurückgeben
    const result = std.mem.toBytes(sum);

    plugin.output(result[0..]);
    return 0;
}
