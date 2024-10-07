export namespace rimage {
	
	export class ImageOptions {
	    dest_format: number;
	    resample_filter: number;
	    desc_path: string;
	    dest_width: number;
	    dest_height: number;
	    jpeg_quality: number;
	    gif_number_of_colors: number;
	    tiff_compression: number;
	    png_compression: number;
	    auto_orientation: boolean;
	    cpu_memory_usage: number;
	
	    static createFrom(source: any = {}) {
	        return new ImageOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.dest_format = source["dest_format"];
	        this.resample_filter = source["resample_filter"];
	        this.desc_path = source["desc_path"];
	        this.dest_width = source["dest_width"];
	        this.dest_height = source["dest_height"];
	        this.jpeg_quality = source["jpeg_quality"];
	        this.gif_number_of_colors = source["gif_number_of_colors"];
	        this.tiff_compression = source["tiff_compression"];
	        this.png_compression = source["png_compression"];
	        this.auto_orientation = source["auto_orientation"];
	        this.cpu_memory_usage = source["cpu_memory_usage"];
	    }
	}
	export class SupportedOutputImageType {
	    name: string;
	    value: number;
	
	    static createFrom(source: any = {}) {
	        return new SupportedOutputImageType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.value = source["value"];
	    }
	}
	export class SupportedResampleFilterType {
	    name: string;
	    value: number;
	
	    static createFrom(source: any = {}) {
	        return new SupportedResampleFilterType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.value = source["value"];
	    }
	}

}

