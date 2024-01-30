import { Injectable } from '@nestjs/common';
import { CreateOrderDto } from './dto/create-order.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { Order } from './entities/order.entity';
import { In, Repository } from 'typeorm';
import { Product } from 'src/products/entities/product.entity';

@Injectable()
export class OrdersService {
  constructor(
    @InjectRepository(Order) private orderRepository: Repository<Order>,
    @InjectRepository(Product) private productRepository: Repository<Product>,
  ) {}

  async create(createOrderDto: CreateOrderDto & { client_id: number }) {
    const productsIds = createOrderDto.items.map((item) => item.product_id);

    const uniqueProductIds = [...new Set(productsIds)];

    const products = await this.productRepository.findBy({
      id: In(uniqueProductIds),
    });

    if (products.length !== uniqueProductIds.length) {
      throw new Error(
        `some product do not exists. products sent ${productsIds}, products found ${products.map((product) => product.id)}`,
      );
    }

    const order = Order.create({
      client_id: createOrderDto.client_id,
      items: createOrderDto.items.map((item) => {
        const product = products.find(
          (product) => product.id === item.product_id,
        );

        return {
          price: product.price,
          quantity: item.quantity,
          product_id: item.product_id,
        };
      }),
    });

    await this.orderRepository.save(order);

    return order;
  }

  findAll(client_id: number) {
    return this.orderRepository.find({
      where: {
        client_id,
      },
      order: {
        creadted_at: 'DESC',
      },
    });
  }

  findOne(id: string, client_id: number) {
    return this.orderRepository.findOneByOrFail({
      id,
      client_id,
    });
  }
}
